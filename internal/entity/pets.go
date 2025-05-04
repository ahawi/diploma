package entity

type Pet struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Gender      string  `db:"gender"`
	Breed       string  `db:"breed"`
	WoolLength  float64 `db:"wool_length"`
	Color       string  `db:"color"`
	Personality string  `db:"personality"`
	Size        float64 `db:"size"`
	Age         float64 `db:"age"`
	Weight      float64 `db:"weight"`
	Health      string  `db:"health"`
	Experience  string  `db:"experience"`
}

func (p *Pet) ToFeatures() *PetFeatures {
	if p == nil {
		return nil
	}
	res := &PetFeatures{
		ID: p.ID,
		Features: map[string]float64{
			"WoolLength": p.WoolLength,
			"Size":       p.Size,
			"Age":        p.Age,
			"Weight":     p.Weight,
		},
	}

	return res
}

func (p *Pet) ToShort() *PetShort {
	if p == nil {
		return nil
	}
	return &PetShort{
		ID:     p.ID,
		Name:   p.Name,
		Breed:  p.Breed,
		Gender: p.Gender,
		Age:    p.Age,
	}
}

type Pets []*Pet

func (p Pets) ToShort() []*PetShort {
	res := make([]*PetShort, len(p))
	for i, x := range p {
		res[i] = x.ToShort()
	}
	return res
}

type PetShort struct {
	ID     int64   `db:"id"`
	Name   string  `db:"name"`
	Breed  string  `db:"breed"`
	Gender string  `db:"gender"`
	Age    float64 `db:"age"`
}

type PetsShort []*PetShort

func (ps PetsShort) IDs() []int64 {
	res := make([]int64, len(ps))
	for i, x := range ps {
		res[i] = x.ID
	}
	return res
}

type PetFeatures struct {
	ID       int64
	Features map[string]float64
}

type PetsFeatures []*PetFeatures

func (pf PetsFeatures) IDs() []int64 {
	res := make([]int64, len(pf))
	for i, x := range pf {
		res[i] = x.ID
	}
	return res
}

type PetFilter struct {
	Gender *string
	Breed  *string
	Size   *float64
	Weight *float64
	Age    *float64
}
