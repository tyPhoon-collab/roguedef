package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*ebiten.Image
	*Transform
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)
	op.GeoM.Translate(s.Pos.X, s.Pos.Y)
	op.GeoM.Rotate(s.Rot)
	screen.DrawImage(s.Image, op)
}

func (s *Sprite) WithTransform(transform *Transform) *Sprite {
	if transform == nil {
		transform = NewTransform()
	}
	s.Transform = transform
	return s
}

func NewSprite(image *ebiten.Image) *Sprite {
	return (&Sprite{
		Image: image,
	}).WithTransform(nil)
}
