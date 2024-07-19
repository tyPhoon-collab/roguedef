package domain

type AttackStatus struct {
	Damage int
}

func NewAttackStatus(damage int) *AttackStatus {
	return &AttackStatus{
		Damage: damage,
	}
}
