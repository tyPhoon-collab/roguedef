package main

import (
	"log"
	"roguedef/object"
	"roguedef/rect"
	"roguedef/system"
	"roguedef/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 = vector.Vec2

func main() {
	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")

	game := system.NewGame()

	player := object.NewPlayer(Vec2{X: 160, Y: 590})
	cursor := object.NewCursor()
	bulletSpawner := object.NewBulletSpawner(player)
	enemySpawner := object.NewEnemySpawner(rect.Rect{
		Min: Vec2{X: 20, Y: 0},
		Max: Vec2{X: 300, Y: 10},
	}).WithPlayer(player)
	levelManager := object.NewLevelManager(enemySpawner)
	debug := object.NewDebug()

	game.AddObjectWithData(player).WithTag("player")
	game.AddObjectWithData(cursor)
	game.AddObjectWithData(bulletSpawner).WithTag("bullet_spawner")
	game.AddObjectWithData(enemySpawner)
	game.AddObjectWithData(levelManager).WithTag("level_manager")
	game.AddObjectWithData(debug)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
