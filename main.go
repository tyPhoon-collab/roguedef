package main

import (
	"log"
	"roguedef/object/game"
	"roguedef/object/title"
	"roguedef/rect"
	"roguedef/system"
	"roguedef/vector"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 = vector.Vec2

func main() {
	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")

	routes := system.Routes{
		"game":  buildGame,
		"title": buildTitle,
	}

	if err := ebiten.RunGame(system.NewScene(routes, "title")); err != nil {
		log.Fatal(err)
	}
}

func buildTitle(scene *system.Scene) ebiten.Game {
	g := system.NewGame()

	ui := title.NewUI(scene)

	g.AddObjectWithData(ui).WithTag("ui")

	return g
}

func buildGame(scene *system.Scene) ebiten.Game {
	g := system.NewGame()

	background := game.NewBackground()
	ui := game.NewUI(scene)
	player := game.NewPlayer(Vec2{X: 160, Y: 590})
	bulletSpawner := game.NewBulletSpawner(player)
	enemySpawner := game.NewEnemySpawner(rect.Rect{
		Min: Vec2{X: 20, Y: 0},
		Max: Vec2{X: 300, Y: 10},
	}, time.Second)
	phaseManager := game.NewPhaseManager(enemySpawner)
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
