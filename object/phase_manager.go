package object

import (
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PhaseManager struct {
	enemySpawner *EnemySpawner
	game         *Game
	level        int
}

func (l *PhaseManager) Register(g *Game, o *system.Object) {
	l.game = g
}

func (l *PhaseManager) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		l.NextLevel()
	}
}

func (l *PhaseManager) NextLevel() {
	l.SetLevel(l.level + 1)
}

func (l *PhaseManager) SetLevel(level int) {
	l.level = level
	l.enemySpawner.SetFrequency(time.Duration(5000.0/(l.level*10)+100) * time.Millisecond)
}

func NewPhaseManager(enemySpawner *EnemySpawner) *PhaseManager {
	m := &PhaseManager{
		enemySpawner: enemySpawner,
	}
	m.SetLevel(1)

	return m
}
