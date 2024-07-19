package domain

type Status struct {
	Hp int
}

func NewStatus(hp int) *Status {
	return &Status{
		Hp: hp,
	}
}
