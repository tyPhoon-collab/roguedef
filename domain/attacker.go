package domain

type Attacker interface {
	Attack(status *Status)
}
