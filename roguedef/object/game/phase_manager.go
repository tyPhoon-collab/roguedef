package game

import (
	"roguedef/system"
)

type PhaseManager struct {
	phase int
}

func (l *PhaseManager) Register(g *Game, o *system.Object) {}

func (l *PhaseManager) Phase() int {
	return l.phase
}

func (l *PhaseManager) NextPhase() {
	l.phase++
}

func NewPhaseManager() *PhaseManager {
	m := &PhaseManager{
		phase: 1,
	}

	return m
}
