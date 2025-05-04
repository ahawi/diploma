package entity

type Form struct {
	Feedbacks []Feedback
}

type Feedback struct {
	UserID int64
	PetID  int64
	Score  float64
}

type Feedbacks []*Feedback

func (f Feedbacks) ByUserID() map[int64]*Feedback {
	if f == nil {
		return map[int64]*Feedback{}
	}

	m := make(map[int64]*Feedback)
	for _, x := range f {
		m[x.UserID] = x
	}

	return m
}

func (f Feedbacks) ByPetID() map[int64]*Feedback {
	if f == nil {
		return map[int64]*Feedback{}
	}

	m := make(map[int64]*Feedback)
	for _, x := range f {
		m[x.PetID] = x
	}

	return m
}
