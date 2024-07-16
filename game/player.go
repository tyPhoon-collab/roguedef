package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/trait"
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type Player struct {
	*trait.Sprite
	intersect trait.Intersector
}

func (p *Player) Intersect() trait.Intersector {
	return p.intersect
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
		p.intersect.Trans().Move(delta)
	}
}

func (p *Player) OnIntersect(other trait.Intersector) {
	fmt.Println("Intersect. Me:", p, " Other:", other)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Sprite.Draw(screen)
}

func (p *Player) String() string {
	return "Player: " + p.Pos.String()
}

func NewPlayer(pos Vec2) (*Player, error) {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		return nil, err
	}

	transform := trait.NewTransform()

	transform.MoveTo(pos)

	return &Player{
		Sprite:    trait.NewSprite(playerImage).WithTransform(transform),
		intersect: trait.NewCircle().WithTransform(transform).FromImage(playerImage),
	}, nil
}
