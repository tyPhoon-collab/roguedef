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
			Exp:   7,
			Speed: 2,
		}
	}

	panic("unknown enemy type")
}

func EnemyTypesByPhase(phase int) []EnemyType {
	len := phase + 10
	types := make([]EnemyType, len)

	switch phase % 2 {
	case 0:
		for i := 0; i < len; i++ {
			types[i] = EnemyTypeSquare
		}
	case 1:
		for i := 0; i < len; i++ {
			types[i] = EnemyTypeTriangle
		}
	default:
		panic("unknown phase")
	}

	return types
}
