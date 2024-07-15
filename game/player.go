package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/trait"
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type Player struct {
	*trait.Sprite
	trait.Intersector
}

func (p *Player) Update() {
	dir := Vec2{}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dir.X = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dir.X = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dir.Y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dir.Y = 1
	}

	p.move(dir)
}

func (p *Player) move(dir Vec2) {
	if dir, err := dir.DivScalar(10); err == nil {
		delta := dir.MulScalar(10)
		p.Move(delta)
		p.Intersector.Trans().Move(delta)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Sprite.Draw(screen)
	p.Intersector.Draw(screen)
}

func NewPlayer() (*Player, error) {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		return nil, err
	}

	transform := trait.NewTransform()

	sprite := trait.NewSpriteWithTransform(playerImage, transform)
	circle := trait.NewCircleFromImage(playerImage)

	circle.Move(transform.Pos)

	return &Player{
		Sprite:      sprite,
		Intersector: circle,
	}, nil
}
