package domain

type Status struct {
	Hp int
}

type EnemyStatus struct {
	Status

	Exp   int
	Speed float64
}
