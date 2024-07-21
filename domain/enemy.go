package domain

import "math"

func ModifyStatusByPhase(phase int, status *EnemyStatus) {
	weight := float64(phase-1)/10 + 1
	status.Hp = int(math.Round(float64(status.Hp) * weight))
}

type EnemyType int

const (
	EnemyTypeSquare EnemyType = iota
	EnemyTypeTriangle
)

func (t EnemyType) Status() EnemyStatus {
	switch t {
	case EnemyTypeSquare:
		return EnemyStatus{
			Status: Status{
				Hp: 10,
			},
			Exp:   10,
			Speed: 1,
		}
	case EnemyTypeTriangle:
		return EnemyStatus{
			Status: Status{
				Hp: 5,
			},
			Exp:   5,
			Speed: 2,
		}
	}

	panic("unknown enemy type")
}

func EnemyBlueprintByPhase(phase int) (EnemyType, EnemyStatus) {
	if phase%2 == 1 {
		return EnemyTypeSquare, EnemyTypeSquare.Status()
	} else {
		return EnemyTypeTriangle, EnemyTypeTriangle.Status()
	}
}
