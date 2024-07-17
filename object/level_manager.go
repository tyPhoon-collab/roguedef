package object

import (
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LevelManager struct {
	enemySpawner *EnemySpawner
	game         *Game
	level        int
}

func (l *LevelManager) Register(g *Game, o *system.Object) {
	l.game = g
}

func (l *LevelManager) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		l.NextLevel()
	}
}

func (l *LevelManager) NextLevel() {
	l.SetLevel(l.level + 1)
}

func (l *LevelManager) SetLevel(level int) {
	l.level = level
	l.enemySpawner.SetFrequency(time.Duration(5000.0/(l.level*10)+100) * time.Millisecond)
}

func NewLevelManager(enemySpawner *EnemySpawner) *LevelManager {
	m := &LevelManager{
		enemySpawner: enemySpawner,
	}
	m.SetLevel(1)

	return m
}
