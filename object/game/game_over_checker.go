package game

import "roguedef/system"

type GameOverChecker struct {
	game *Game
	ui   *UI
}

func (c *GameOverChecker) Register(g *Game, o *system.Object) {
	c.game = g
	c.ui = g.ObjectByTag("ui").Data.(*UI)
}

func (c *GameOverChecker) Update() {
	for o := range c.game.ObjectsByTag("enemy") {
		if enemy, ok := o.Data.(*Enemy); ok {
			if c.game.IsOutside(enemy.Pos) {
				go c.ui.WaitShowGameOver()
			}
		}
	}
}

func NewGameOverChecker() *GameOverChecker {
	return &GameOverChecker{}
}
