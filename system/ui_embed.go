package system

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type UIEmbed struct {
	ui          *ebitenui.UI
	face        font.Face
	buttonImage *widget.ButtonImage
}

func (u *UIEmbed) BasicButtonOpts(text string, do func(args *widget.ButtonClickedEventArgs)) widget.ButtonOpt {
	buttonTextColor := &widget.ButtonTextColor{
		Idle: color.RGBA{255, 255, 255, 255},
	}
	return func(b *widget.Button) {
		opts := []widget.ButtonOpt{
			widget.ButtonOpts.Text(text, u.face, buttonTextColor),
			widget.ButtonOpts.TextPadding(widget.Insets{Left: 20, Right: 20, Top: 5, Bottom: 5}),
			widget.ButtonOpts.Image(u.buttonImage),
			widget.ButtonOpts.ClickedHandler(do),
		}

		for _, opt := range opts {
			opt(b)
		}
	}
}

func (u *UIEmbed) UI() *ebitenui.UI {
	return u.ui
}

func (u *UIEmbed) Draw(screen *ebiten.Image) {
	u.ui.Draw(screen)
}

func (u *UIEmbed) Priority() int {
	return 1000
}

func (u *UIEmbed) Update() {
	u.ui.Update()
}

func NewUIEmbed(container *widget.Container) *UIEmbed {
	face, err := LoadFont()
	if err != nil {
		panic(err)
	}

	buttonImage, err := LoadButtonImage()
	if err != nil {
		panic(err)
	}

	ui := &ebitenui.UI{Container: container}

	return &UIEmbed{ui: ui, face: face, buttonImage: buttonImage}
}

func LoadButtonImage() (*widget.ButtonImage, error) {
	img, _, err := ebitenutil.NewImageFromFile("resources/images/ui/button.png")

	if err != nil {
		return nil, err
	}
	idle := image.NewNineSliceSimple(img, 16, 32)

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   idle,
		Pressed: idle,
	}, nil
}

func LoadFont() (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
