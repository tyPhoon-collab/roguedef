package roguedef

import (
	"roguedef/ds"
	"roguedef/rect"
	"roguedef/roguedef/object/equip"
	"roguedef/roguedef/object/game"
	"roguedef/roguedef/object/title"
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 = ds.Vec2

func NewRogueDef() ebiten.Game {
	routes := system.Routes{
		"game":  buildGame,
		"title": buildTitle,
		"equip": buildEquip,
	}

	return system.NewScene(routes, "title")
}

func buildTitle(scene *system.Scene) *system.Game {
	g := system.NewGame()

	g.AddObjectWithData(title.NewUI(scene)).WithTag("ui")
	g.AddObjectWithData(title.NewBackground(320, 640))

	return g
}

func buildEquip(scene *system.Scene) *system.Game {
	g := system.NewGame()

	g.AddObjectWithData(equip.NewBackground(320, 640))
	g.AddObjectWithData(equip.NewUI(scene)).WithTag("ui")

	return g
}

func buildGame(scene *system.Scene) *system.Game {
	g := system.NewGame()

	background := game.NewBackground()
	ui := game.NewUI(scene)
	player := game.NewPlayer(Vec2{X: 160, Y: 600})
	bulletSpawner := game.NewBulletSpawner(Vec2{X: 160, Y: 560})
	enemySpawner := game.NewEnemySpawner(rect.Rect{
		Min: Vec2{X: 20, Y: 0},
		Max: Vec2{X: 300, Y: 10},
	}, 2*time.Second)
	phaseManager := game.NewPhaseManager()
	gameOverChecker := game.NewGameOverChecker()
	debug := game.NewDebug()

	g.AddObjectWithData(background)
	g.AddObjectWithData(ui).WithTag("ui")
	g.AddObjectWithData(player).WithTag("player")
	g.AddObjectWithData(bulletSpawner).WithTag("bullet_spawner")
	g.AddObjectWithData(enemySpawner).WithTag("enemy_spawner")
	g.AddObjectWithData(phaseManager).WithTag("phase_manager")
	g.AddObjectWithData(gameOverChecker)
	g.AddObjectWithData(debug)

	return g
}
