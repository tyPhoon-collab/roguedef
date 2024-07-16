package main

import (
	"log"
	"roguedef/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
