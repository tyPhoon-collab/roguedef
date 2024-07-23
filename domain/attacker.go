package domain

type AttackedContext struct {
	AppliedDamage int
}

type Attacker interface {
	Attack(status *Status) AttackedContext
}
