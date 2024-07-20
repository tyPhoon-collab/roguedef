package system

import (
	"fmt"
	"roguedef/ds"

	"github.com/hajimehoshi/ebiten/v2"
)

type context struct {
	route string
	game  ebiten.Game
}

type Scene struct {
	routes Routes
	stack  *ds.Stack[context]
}

func (s *Scene) Current() context {
	c, ok := s.stack.Peek()
	if !ok {
		panic("stack is empty")
	}
	return c
}

func (s *Scene) Push(route string) error {
	s.push(route)
	return nil
}

func (s *Scene) Reload() error {
	c, _ := s.stack.Pop()
	return s.Push(c.route)
}

func (s *Scene) Pop() bool {
	if s.stack.Len() <= 1 {
		return false
	}

	_, _ = s.stack.Pop()
	return true
}

func (s *Scene) Update() error {
	err := s.Current().game.Update()
	return err
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.Current().game.Draw(screen)
}

func (s *Scene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.Current().game.Layout(outsideWidth, outsideHeight)
}

func (s *Scene) buildGame(route string) (ebiten.Game, error) {
	builder, ok := s.routes[route]
	if !ok {
		return nil, fmt.Errorf("route not found: %s", route)
	}
	return builder(s), nil
}

func (s *Scene) push(route string) error {
	game, err := s.buildGame(route)

	if err != nil {
		return err
	}

	c := context{
		route: route,
		game:  game,
	}
	s.stack.Push(c)

	return nil
}

func NewScene(routes Routes, initRoute string) *Scene {
	s := &Scene{
		routes: routes,
		stack:  ds.NewStack[context](),
	}

	if err := s.Push(initRoute); err != nil {
		panic(err)
	}

	return s
}
