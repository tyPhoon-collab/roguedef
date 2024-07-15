package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player            *Player
	drawers           []trait.Drawer
	updaters          []trait.Updater
	intersects        []trait.Intersector
	intersectHandlers map[trait.Intersector]trait.IntersectHandler
}

func (g *Game) Update() error {
	for _, o := range g.updaters {
		o.Update()
	}

	for i := 0; i < len(g.intersects); i++ {
		for j := i + 1; j < len(g.intersects); j++ {
			o1 := g.intersects[i]
			o2 := g.intersects[j]
			if o1.Intersects(o2) {
				if handler, ok := g.intersectHandlers[o1]; ok {
					handler.OnIntersect(o2)
				}
				if handler, ok := g.intersectHandlers[o2]; ok {
					handler.OnIntersect(o1)
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.player.Pos.String())

	for _, o := range g.drawers {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func NewGame() *Game {
	player, err := NewPlayer()
	cursor := NewCursor()

	if err != nil {
		panic(err)
	}

	return &Game{
		player:     player,
		drawers:    []trait.Drawer{player, cursor},
		updaters:   []trait.Updater{player, cursor},
		intersects: []trait.Intersector{player.Intersector, cursor.Intersector},
		intersectHandlers: map[trait.Intersector]trait.IntersectHandler{
			player.Intersector: player,
		},
	}
}
