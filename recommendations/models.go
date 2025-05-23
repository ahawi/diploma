package models

type Dog struct {
	ID         int
	Breed      string
	Size       string
	Age        int
	Activity   string
	Temperament string
}

type User struct {
	ID              int
	PreferredBreeds []string
	PreferredSize   string
	AgeRange        [2]int
	ActivityLevel   string
	Temperament     string
}

type Recommendation struct {
	UserID int
	DogIDs []int
}
