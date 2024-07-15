package game

import (
	"fmt"
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	*Player
	drawers    []trait.Drawer
	updaters   []trait.Updater
	intersects []trait.Intersector
}

func (g *Game) Update() error {
	for _, o := range g.updaters {
		o.Update()
	}

	for _, o := range g.intersects {
		for _, other := range g.intersects {
			if o == other {
				continue
			}
			if o.Intersects(other) {
				fmt.Println(o, other)
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.Player.Pos.String())

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
		Player:     player,
		drawers:    []trait.Drawer{player, cursor},
		updaters:   []trait.Updater{player, cursor},
		intersects: []trait.Intersector{player.Intersector, cursor.Intersector},
	}
}
