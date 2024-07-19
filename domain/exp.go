package domain

import (
	"math"
)

type ExpManager struct {
	totalExp int
	exp      int
	level    int
}

func (m *ExpManager) Exp() int {
	return m.exp
}

func (m *ExpManager) Level() int {
	return m.level
}

// return true if level up
func (m *ExpManager) AddExp(exp int) bool {
	m.exp += exp
	m.totalExp += exp
	return m.levelUp()
}

// return true if level up
func (m *ExpManager) levelUp() bool {
	if m.exp >= ExpToNextLevel(m.level) {
		m.level++
		m.exp -= ExpToNextLevel(m.level - 1)

		return true
	}

	return false
}

func ExpToNextLevel(level int) int {
	base := 100
	multiplier := 1.15
	return base * int(math.Pow(multiplier, float64(level)))
}

func NewExpManager() *ExpManager {
	return &ExpManager{
		level: 1,
	}
}
