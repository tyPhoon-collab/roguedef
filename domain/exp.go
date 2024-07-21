package domain

import (
	"math"
	"roguedef/system"
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

		for m.levelUp() {
			// do nothing
		}

		return true
	}

	return false
}

func ExpToNextLevel(level int) int {
	base := 100
	multiplier := 1.15
	return system.MulIntByFloat(base, math.Pow(multiplier, float64(level-1)))
}

func NewExpManager() *ExpManager {
	return &ExpManager{
		level: 1,
	}
}
