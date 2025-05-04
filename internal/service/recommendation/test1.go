package recommendation

import (
	"math"
	"sort"
)

type Item struct {
	ID       int
	Features map[string]float64
}

type User struct {
	ID      int
	Ratings map[int]float64 // itemID: rating
}

type Test1S struct {
	Users               map[int]User
	Items               map[int]Item
	ContentWeight       float64
	CollaborativeWeight float64
}

func NewHybridRecommender(contentWeight, collaborativeWeight float64) *Test1S {
	return &Test1S{
		Users:               make(map[int]User),
		Items:               make(map[int]Item),
		ContentWeight:       contentWeight,
		CollaborativeWeight: collaborativeWeight,
	}
}

func (hr *Test1S) collaborativeScore(targetUserID int, itemID int) float64 {
	targetUser := hr.Users[targetUserID]
	sumSimilarity := 0.0
	sumRating := 0.0

	for _, user := range hr.Users {
		if user.ID == targetUserID {
			continue
		}

		if rating, exists := user.Ratings[itemID]; exists {
			similarity := hr.userSimilarity(targetUser, user)
			sumSimilarity += similarity
			sumRating += similarity * rating
		}
	}

	if sumSimilarity == 0 {
		return 0.0
	}
	return sumRating / sumSimilarity
}

func (hr *Test1S) userSimilarity(a, b User) float64 {
	dotProduct := 0.0
	magA := 0.0
	magB := 0.0

	for itemID, aRating := range a.Ratings {
		if bRating, exists := b.Ratings[itemID]; exists {
			dotProduct += aRating * bRating
			magA += aRating * aRating
			magB += bRating * bRating
		}
	}

	if magA == 0 || magB == 0 {
		return 0.0
	}
	return dotProduct / (math.Sqrt(magA) * math.Sqrt(magB))
}

func (hr *Test1S) contentScore(targetUserID int, itemID int) float64 {
	targetUser := hr.Users[targetUserID]
	item := hr.Items[itemID]

	userFeatures := make(map[string]float64)
	for _, rating := range targetUser.Ratings {
		for feature, value := range hr.Items[itemID].Features {
			userFeatures[feature] += value * rating
		}
	}

	sum := 0.0
	for _, val := range userFeatures {
		sum += val * val
	}
	norm := math.Sqrt(sum)
	if norm != 0 {
		for f := range userFeatures {
			userFeatures[f] /= norm
		}
	}

	sum = 0.0
	for _, val := range item.Features {
		sum += val * val
	}
	itemNorm := math.Sqrt(sum)

	dotProduct := 0.0
	for feature, userVal := range userFeatures {
		itemVal := item.Features[feature]
		if itemNorm != 0 {
			itemVal /= itemNorm
		}
		dotProduct += userVal * itemVal
	}

	return dotProduct
}

func (hr *Test1S) HybridScore(targetUserID, itemID int) float64 {
	collabScore := hr.collaborativeScore(targetUserID, itemID)
	contentScore := hr.contentScore(targetUserID, itemID)

	return hr.CollaborativeWeight*collabScore + hr.ContentWeight*contentScore
}

func (hr *Test1S) Recommend(targetUserID int, topN int) []int {
	targetUser := hr.Users[targetUserID]
	itemScores := make(map[int]float64)

	for itemID := range hr.Items {
		if _, exists := targetUser.Ratings[itemID]; !exists {
			itemScores[itemID] = hr.HybridScore(targetUserID, itemID)
		}
	}

	type kv struct {
		Key   int
		Value float64
	}

	var ss []kv
	for k, v := range itemScores {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var result []int
	for i := 0; i < topN && i < len(ss); i++ {
		result = append(result, ss[i].Key)
	}

	return result
}

// func main() {
// 	recommender := NewHybridRecommender(0.5, 0.5)

// 	recommender.Items = map[int]Item{
// 		1: {ID: 1, Features: map[string]float64{"action": 2.5, "comedy": 1.2}},
// 		2: {ID: 2, Features: map[string]float64{"comedy": 2.1, "drama": 1.7}},
// 		3: {ID: 3, Features: map[string]float64{"action": 1.8, "sci-fi": 2.5}},
// 		4: {ID: 4, Features: map[string]float64{"drama": 2.0, "romance": 1.5}},
// 	}

// 	recommender.Users = map[int]User{
// 		1: {
// 			ID: 1,
// 			Ratings: map[int]float64{
// 				1: 5.0,
// 				2: 4.5,
// 			},
// 		},
// 		2: {
// 			ID: 2,
// 			Ratings: map[int]float64{
// 				3: 4.0,
// 				4: 3.5,
// 			},
// 		},
// 	}

// 	recommendations := recommender.Recommend(1, 3)

// 	fmt.Println("Рекомендованные товары:")
// 	for i, itemID := range recommendations {
// 		fmt.Printf("%d. Товар %d (Счет: %.2f)\n", i+1, itemID,
// 			recommender.HybridScore(1, itemID))
// 	}
// }
