package game

import "roguedef/system"

type GameOverChecker struct {
	game     *Game
	ui       *UI
	gameOver bool
}

func (c *GameOverChecker) Register(g *Game, o *system.Object) {
	c.game = g
	c.ui = g.ObjectByTag("ui").Data.(*UI)
}

func (c *GameOverChecker) Update() {
	if !c.ShouldCheck() {
		return
	}

	for o := range c.game.ObjectsByTag("enemy") {
		if enemy, ok := o.Data.(*Enemy); ok {
			if c.game.IsOutside(enemy.Pos) {
				c.gameOver = true
				go c.ui.WaitShowGameOver()
			}
		}
	}
}

func (c *GameOverChecker) ShouldCheck() bool {
	return !c.gameOver && system.TimeScale != 0
}

func NewGameOverChecker() *GameOverChecker {
	return &GameOverChecker{}
}
