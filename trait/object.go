package trait

import "github.com/hajimehoshi/ebiten/v2"

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Updater interface {
	Update()
}

type Object interface {
	Updater
	Drawer
}
