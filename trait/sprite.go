package trait

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*ebiten.Image
	*Transform
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.Pos.X, s.Pos.Y)
	op.GeoM.Rotate(s.Rot)
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)
	screen.DrawImage(s.Image, op)
}

func NewSprite(image *ebiten.Image) *Sprite {
	return NewSpriteWithTransform(image, nil)
}

func NewSpriteWithTransform(image *ebiten.Image, transform *Transform) *Sprite {
	if transform == nil {
		transform = NewTransform()
	}
	return &Sprite{
		Image:     image,
		Transform: transform,
	}
}
