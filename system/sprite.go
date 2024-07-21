package system

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Alignment int

const (
	AlignTopLeft Alignment = iota
	AlignCenter
)

type Sprite struct {
	*ebiten.Image
	*Transform
	alignment Alignment
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)

	switch s.alignment {
	case AlignCenter:
		op.GeoM.Translate(-float64(s.Image.Bounds().Dx())*s.Scale.X/2, -float64(s.Image.Bounds().Dy())*s.Scale.Y/2)
	}

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

func (s *Sprite) WithAlignment(alignment Alignment) *Sprite {
	s.alignment = alignment
	return s
}

func NewEmptySprite() *Sprite {
	img := ebiten.NewImage(1, 1)
	return (&Sprite{Image: img}).WithTransform(nil).WithAlignment(AlignCenter)
}

func NewSprite(image *ebiten.Image) *Sprite {
	return (&Sprite{
		Image: image,
	}).WithTransform(nil).WithAlignment(AlignCenter)
}

func NewColorSprite(width int, height int, color color.Color) *Sprite {
	img := ebiten.NewImage(width, height)
	img.Fill(color)

	return (&Sprite{Image: img}).WithTransform(nil).WithAlignment(AlignCenter)
}
