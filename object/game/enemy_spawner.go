package game

import (
	"roguedef/domain"
	"roguedef/ds"
	"roguedef/resources"
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game         *Game
	spawnRange   Rect
	player       *Player
	phaseManager *PhaseManager
	queue        *ds.Queue[domain.EnemyType]
	*system.Looper
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
	s.player = g.ObjectByTag("player").Data.(*Player)
	s.phaseManager = g.ObjectByTag("phase_manager").Data.(*PhaseManager)

	s.buildQueue()
}

func (s *EnemySpawner) Update() {
	s.Looper.Update()
}

func (s *EnemySpawner) buildQueue() {
	phase := s.phaseManager.phase
	types := domain.EnemyTypesByPhase(phase)
	s.queue = ds.NewQueueFrom(types)
}

func (s *EnemySpawner) addEnemy() {
	if s.queue.IsEmpty() {
		s.phaseManager.NextPhase()
		system.ScaleDuration(&s.Frequency, 0.85)

		s.buildQueue()
	}

	t, _ := s.queue.Pop()
	phase := s.phaseManager.phase
	status := t.Status()

	domain.ModifyStatusByPhase(phase, &status)

	var enemy *Enemy

	switch t {
	case domain.EnemyTypeSquare:
		enemy = NewEnemySquare(24)
	case domain.EnemyTypeTriangle:
		enemy = NewEnemyTriangle(24)
	case domain.EnemyTypeBoss:
		enemy = NewEnemyTriangle(64)
	default:
		enemy = NewEnemyFromResource(resources.GopherImage)
	}

	enemy = enemy.
		WithPlayer(s.player).
		WithStatus(&status)

	enemy.Pos = s.spawnRange.RandomPoint()
	s.game.AddObjectWithData(enemy).WithTag("enemy")
}

func NewEnemySpawner(spawnRange Rect, frequency time.Duration) *EnemySpawner {
	s := &EnemySpawner{
		spawnRange: spawnRange,
		queue:      ds.NewQueue[domain.EnemyType](),
	}
	s.Looper = system.NewLooper(frequency, s.addEnemy)

	return s
}
