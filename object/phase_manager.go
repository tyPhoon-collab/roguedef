package object

import (
	"roguedef/system"
	"time"
)

type PhaseManager struct {
	enemySpawner *EnemySpawner
	game         *Game
	phase        int
	*system.Looper
}

func (l *PhaseManager) Register(g *Game, o *system.Object) {
	l.game = g
}

func (l *PhaseManager) Update() {
	l.Looper.Update()
}

func (l *PhaseManager) NextPhase() {
	l.phase++
	system.ScaleDuration(&l.enemySpawner.Frequency, 0.9)
}

func NewPhaseManager(enemySpawner *EnemySpawner) *PhaseManager {
	m := &PhaseManager{
		enemySpawner: enemySpawner,
		phase:        1,
	}

	m.Looper = system.NewLooper(10*time.Second, m.NextPhase)

	return m
}
