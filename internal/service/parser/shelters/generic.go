package shelters

import (
	"net/http"
	"strings"

	"github.com/ahawi/diploma/internal/entity"

	"github.com/PuerkitoBio/goquery"
)

func ParseGeneric(shelter entity.Shelter) ([]entity.Dog, error) {
	res, err := http.Get(shelter.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var dogs []entity.Dog
	doc.Find(".dog-card").Each(func(i int, s *goquery.Selection) {
		dog := entity.Dog{
			Name:     strings.TrimSpace(s.Find(".dog-name").Text()),
			Age:      strings.TrimSpace(s.Find(".dog-age").Text()),
			Breed:    strings.TrimSpace(s.Find(".dog-breed").Text()),
			Gender:   strings.TrimSpace(s.Find(".dog-gender").Text()),
			ImageURL: s.Find("img").AttrOr("src", ""),
		}
		dogs = append(dogs, dog)
	})

	return dogs, nil
}
