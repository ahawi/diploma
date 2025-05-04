package recommendation

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/render"
)

type Cache struct {
	data sync.Map
	ttl  time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		data: sync.Map{},
		ttl:  ttl,
	}
}

func (c *Cache) GetRecommendations(state *UserState, count int) []Item2 {
	if cached, ok := c.data.Load(state.Cursor); ok {
		return cached.([]Item2)
	}
	return nil
}

func (c *Cache) StoreRecommendations(cursor string, items []Item2) {
	c.data.Store(cursor, items)
	time.AfterFunc(c.ttl, func() {
		c.data.Delete(cursor)
	})
}

type Test2S struct {
	mu          sync.Mutex
	userStates  map[string]*UserState
	cache       *Cache
	itemCatalog []Item2
}

type UserState struct {
	LastViewed      []string
	Recommendations []string
	Cursor          string
	LastUpdated     time.Time
}

type Item2 struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type Interaction struct {
	UserID string `json:"user_id"`
	ItemID string `json:"item_id"`
	Type   string `json:"type"`
}

func NewTest2S() *Test2S {
	return &Test2S{
		userStates:  make(map[string]*UserState),
		itemCatalog: loadCatalog(),
		cache:       NewCache(5 * time.Minute),
	}
}

func loadCatalog() []Item2 {
	return []Item2{
		{ID: "1", Title: "Товар 1", Description: "Описание 1", ImageURL: "image1.jpg"},
		{ID: "2", Title: "Товар 2", Description: "Описание 2", ImageURL: "image2.jpg"},
	}
}

func (rs *Test2S) generateInitialRecommendations() []string {
	return []string{"1", "2", "3"}
}

func generateNewCursor() string {
	return time.Now().Format("20060102150405")
}

func getUserId(r *http.Request) string {
	_ = r
	return "user123"
}

func (rs *Test2S) updateUserState(state *UserState, interaction Interaction) {

	state.LastViewed = append(state.LastViewed, interaction.ItemID)

	state.Recommendations = rs.generateUpdatedRecommendations(state)

	state.LastUpdated = time.Now()
}

func (rs *Test2S) refreshRecommendations(state *UserState) {
	state.Recommendations = rs.generateUpdatedRecommendations(state)
}

func (rs *Test2S) generateUpdatedRecommendations(state *UserState) []string {

	popular := rs.generateInitialRecommendations()
	similar := rs.findSimilarItems(state.LastViewed)

	return append(popular[:len(popular)/2], similar[:len(similar)/2]...)
}

func (rs *Test2S) findSimilarItems(itemIDs []string) []string {

	var similar []string
	for _, id := range itemIDs {
		similar = append(similar, "related_"+id)
	}
	return similar
}

func (rs *Test2S) generateRecommendations(state *UserState, count int, cursor string) []Item2 {

	if cached := rs.cache.GetRecommendations(state, count); cached != nil {
		return cached
	}

	var recommendations []Item2
	for _, itemID := range state.Recommendations {
		for _, item := range rs.itemCatalog {
			if item.ID == itemID {
				recommendations = append(recommendations, item)
				break
			}
		}
	}

	rs.cache.StoreRecommendations(cursor, recommendations)

	return recommendations
}

func (rs *Test2S) getUserState(userID string) *UserState {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if state, exists := rs.userStates[userID]; exists {
		return state
	}

	state := &UserState{
		LastViewed:      []string{},
		Recommendations: rs.generateInitialRecommendations(),
		LastUpdated:     time.Now(),
	}
	rs.userStates[userID] = state
	return state
}

func (rs *Test2S) HandleInteraction(w http.ResponseWriter, r *http.Request) {
	var interaction Interaction
	if err := json.NewDecoder(r.Body).Decode(&interaction); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	rs.mu.Lock()
	defer rs.mu.Unlock()

	state := rs.getUserState(interaction.UserID)
	rs.updateUserState(state, interaction)
	rs.refreshRecommendations(state)

	render.NoContent(w, r)
}

func (rs *Test2S) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	userID := getUserId(r)
	pageSize := 10
	cursor := r.URL.Query().Get("cursor")

	state := rs.getUserState(userID)
	recommendations := rs.generateRecommendations(state, pageSize, cursor)

	render.JSON(w, r, map[string]interface{}{
		"items":  recommendations,
		"cursor": generateNewCursor(),
	})
}
