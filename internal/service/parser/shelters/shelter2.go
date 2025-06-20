package shelters

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/ahawi/diploma/internal/entity"
)

func ParseShelter2(shelter entity.Shelter) ([]entity.Dog, error) {
	res, err := http.Get(shelter.URL + "/animals/dogs")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var dogs []entity.Dog
	doc.Find("div.grid-card").Each(func(i int, s *goquery.Selection) {
		dog := entity.Dog{
			Name:     s.Find("h3.name").Text(),
			Age:      s.Find(".details .age").Text(),
			Breed:    s.Find(".details .breed").Text(),
			Gender:   s.Find(".details .gender").Text(),
			ImageURL: s.Find("img").AttrOr("data-src", ""),
		}
		dogs = append(dogs, dog)
	})

	return dogs, nil
}
