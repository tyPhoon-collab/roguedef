package main

import (
	"log"
	"roguedef/ds"
	"roguedef/roguedef"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 = ds.Vec2

func main() {
	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Hello World (Ebitengine Demo)")

	if err := ebiten.RunGame(roguedef.NewRogueDef()); err != nil {
		log.Fatal(err)
	}
}
