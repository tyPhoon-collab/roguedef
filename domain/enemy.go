package domain

import (
	"math"
	"slices"
)

func ModifyStatusByPhase(phase int, status *EnemyStatus) {
	p := float64(phase)
	weight := (p-1)/10 + 1
	status.Hp = int(math.Round(float64(status.Hp) * weight))
	status.Speed = status.Speed + (0.1 * p)
}

type EnemyType int

const (
	EnemyTypeSquare EnemyType = iota
	EnemyTypeTriangle
	EnemyTypeBoss
)

func (t EnemyType) Status() EnemyStatus {
	switch t {
	case EnemyTypeSquare:
		return EnemyStatus{
			Status: Status{
				Hp: 20,
			},
			Exp:   15,
			Speed: 1,
		}
	case EnemyTypeTriangle:
		return EnemyStatus{
			Status: Status{
				Hp: 10,
			},
			Exp:   7,
			Speed: 2,
		}
	case EnemyTypeBoss:
		return EnemyStatus{
			Status: Status{
				Hp: 200,
			},
			Exp:   200,
			Speed: 0.5,
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
			types[i] = EnemyTypeTriangle
		}
	case 1:
		for i := 0; i < len; i++ {
			types[i] = EnemyTypeSquare
		}
	}

	if phase%5 == 0 {
		types = slices.Insert(types, 0, EnemyTypeBoss)
	}

	return types
}
