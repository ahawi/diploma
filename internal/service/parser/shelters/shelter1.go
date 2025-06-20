package shelters

import (
	"net/http"
	"strings"

	"github.com/ahawi/diploma/internal/entity"

	"github.com/PuerkitoBio/goquery"
)

func ParseShelter1(shelter entity.Shelter) ([]entity.Dog, error) {
	res, err := http.Get(shelter.URL + "/dogs")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var dogs []entity.Dog
	doc.Find("table.dog-table tbody tr").Each(func(i int, row *goquery.Selection) {
		cols := row.Find("td")
		if cols.Length() < 5 {
			return
		}
		dog := entity.Dog{
			Name:     strings.TrimSpace(cols.Eq(0).Text()),
			Age:      strings.TrimSpace(cols.Eq(1).Text()),
			Breed:    strings.TrimSpace(cols.Eq(2).Text()),
			Gender:   strings.TrimSpace(cols.Eq(3).Text()),
			ImageURL: cols.Eq(4).Find("img").AttrOr("src", ""),
		}
		dogs = append(dogs, dog)
	})

	return dogs, nil
}
