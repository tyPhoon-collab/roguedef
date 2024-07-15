package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/trait"
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type Player struct {
	image *ebiten.Image

	trait.HasPosition
}

func (p *Player) Update() {
	delta := Vec2{X: 0, Y: 0}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		delta.X = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		delta.X = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		delta.Y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		delta.Y = 1
	}

	if delta, err := delta.DivScalar(10); err == nil {
		p.Move(delta.MulScalar(10))
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.Pos.X, p.Pos.Y)
	screen.DrawImage(p.image, op)
}

func NewPlayer() (*Player, error) {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		return nil, err
	}

	return &Player{
		image: playerImage,
	}, nil
}
