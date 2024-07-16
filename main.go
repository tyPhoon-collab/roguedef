package main

import (
	"log"
	"roguedef/object"
	"roguedef/system"
	"roguedef/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")

	game := system.NewGame()

	player := object.NewPlayer(vector.Vec2{X: 136, Y: 550})
	cursor := object.NewCursor()
	debug := object.NewDebug()
	bulletSpawner := object.NewBulletSpawner(player)
	enemySpawner := object.NewEnemySpawner()

	game.AddObjectWithData(bulletSpawner)
	game.AddObjectWithData(player)
	game.AddObjectWithData(cursor)
	game.AddObjectWithData(debug)
	game.AddObjectWithData(enemySpawner)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
