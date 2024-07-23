package domain

type Status struct {
	Hp  int
	Def int
}

func (s *Status) Apply(a AttackStatus) AttackedContext {
	dmg := a.Damage - s.Def

	if dmg < 0 {
		dmg = 1
	}
	s.Hp -= dmg

	return AttackedContext{AppliedDamage: dmg}
}

type EnemyStatus struct {
	Status

	Exp   int
	Speed float64
}
