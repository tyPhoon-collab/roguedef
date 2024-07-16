package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/system"
)

type Player struct {
	*system.Transform
	sprite    *system.Sprite
	intersect system.Intersector
}

func (p *Player) Register(o *system.Object) {}

func (p *Player) Intersect() system.Intersector {
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
		p.Move(delta)
	}
}

func (p *Player) OnIntersect(other *system.Object) {
	// fmt.Println("Intersect. Me:", p, " Other:", other)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) String() string {
	return "Player: " + p.Pos.String()
}

func NewPlayer(pos Vec2) (*Player, error) {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		return nil, err
	}

	transform := system.NewTransform()

	transform.MoveTo(pos)

	return &Player{
		Transform: transform,
		sprite:    system.NewSprite(playerImage).WithTransform(transform),
		intersect: system.NewCircle().WithTransform(transform).FromImage(playerImage),
	}, nil
}
