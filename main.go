package main

import (
	"log"
	"roguedef/ds"
	"roguedef/object/game"
	"roguedef/object/title"
	"roguedef/rect"
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 = ds.Vec2

const width, height = 320, 640

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")

	routes := system.Routes{
		"game":  buildGame,
		"title": buildTitle,
	}

	if err := ebiten.RunGame(system.NewScene(routes, "title")); err != nil {
		log.Fatal(err)
	}
}

func buildTitle(scene *system.Scene) *system.Game {
	g := system.NewGame()

	g.AddObjectWithData(title.NewUI(scene)).WithTag("ui")
	g.AddObjectWithData(title.NewBackground(width, height))

	return g
}

func buildGame(scene *system.Scene) *system.Game {
	g := system.NewGame()

	background := game.NewBackground()
	ui := game.NewUI(scene)
	player := game.NewPlayer(Vec2{X: 160, Y: 580})
	bulletSpawner := game.NewBulletSpawner(Vec2{X: 160, Y: 510})
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
