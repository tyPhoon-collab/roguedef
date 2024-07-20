package system

import (
	"roguedef/ds"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	routes  Routes
	current ebiten.Game
	stack   *ds.Stack[ebiten.Game]
}

func (s *Scene) Push(route string) {
	s.stack.Push(s.current)
	s.current = s.routes[route]()
}

func (s *Scene) Pop() bool {
	if s.stack.Len() <= 1 {
		return false
	}

	_, _ = s.stack.Pop()
	s.current, _ = s.stack.Peek()
	return true
}

func (s *Scene) Update() error {
	err := s.current.Update()
	return err
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.current.Draw(screen)
}

func (s *Scene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.current.Layout(outsideWidth, outsideHeight)
}

func NewScene(routes Routes, initRoute string) *Scene {
	s := &Scene{
		routes:  routes,
		current: routes[initRoute](),
		stack:   ds.NewStack[ebiten.Game](),
	}

	s.stack.Push(s.current)

	return s
}
