package object

import (
	"roguedef/system"
	"time"
)

type PhaseManager struct {
	enemySpawner *EnemySpawner
	game         *Game
	phase        int
}

func (l *PhaseManager) Register(g *Game, o *system.Object) {
	l.game = g
}

func (l *PhaseManager) Update() {
	frameCount := system.DurationToFrameCount(time.Duration(10) * time.Second)
	if l.game.FrameCount()%frameCount == 0 {
		l.NextPhase()
	}
}

func (l *PhaseManager) NextPhase() {
	l.SetPhase(l.phase + 1)
}

func (l *PhaseManager) SetPhase(phase int) {
	l.phase = phase
	l.enemySpawner.SetFrequency(time.Duration(5000.0/(l.phase*10)+100) * time.Millisecond)
}

func NewPhaseManager(enemySpawner *EnemySpawner) *PhaseManager {
	m := &PhaseManager{
		enemySpawner: enemySpawner,
	}
	m.SetPhase(1)

	return m
}
