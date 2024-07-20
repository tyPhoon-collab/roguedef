package main

import (
	"log"
	"roguedef/object"
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
		"game": buildGame,
	}

	if err := ebiten.RunGame(system.NewScene(routes, "game")); err != nil {
		log.Fatal(err)
	}
}

func buildGame() ebiten.Game {
	game := system.NewGame()

	background := object.NewBackground()
	ui := object.NewUI()
	player := object.NewPlayer(Vec2{X: 160, Y: 590})
	bulletSpawner := object.NewBulletSpawner(player)
	enemySpawner := object.NewEnemySpawner(rect.Rect{
		Min: Vec2{X: 20, Y: 0},
		Max: Vec2{X: 300, Y: 10},
	}, time.Second)
	phaseManager := object.NewPhaseManager(enemySpawner)
	gameOverChecker := object.NewGameOverChecker()
	debug := object.NewDebug()

	game.AddObjectWithData(background)
	game.AddObjectWithData(ui).WithTag("ui")
	game.AddObjectWithData(player).WithTag("player")
	game.AddObjectWithData(bulletSpawner).WithTag("bullet_spawner")
	game.AddObjectWithData(enemySpawner).WithTag("enemy_spawner")
	game.AddObjectWithData(phaseManager).WithTag("phase_manager")
	game.AddObjectWithData(gameOverChecker)
	game.AddObjectWithData(debug)

	return game
}
