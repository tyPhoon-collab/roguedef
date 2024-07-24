package system

import (
	img "image"
	"image/color"
	"roguedef/resources"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type UIEmbed struct {
	ui          *ebitenui.UI
	face        font.Face
	buttonImage *widget.ButtonImage
}

func (u *UIEmbed) BackgroundImage() widget.ContainerOpt {
	img := LoadImage(resources.WindowImage)
	bg := widget.ContainerOpts.BackgroundImage(image.NewNineSliceSimple(img, 12, 8))
	return bg
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

func (u *UIEmbed) BasicTextOpts(text string) widget.TextOpt {
	return widget.TextOpts.Text(text, u.face, color.White)
}

func (u *UIEmbed) UI() *ebitenui.UI {
	return u.ui
}

func (u *UIEmbed) Container() *widget.Container {
	return u.ui.Container
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

	buttonImage := LoadButtonImage()

	ui := &ebitenui.UI{Container: container}

	return &UIEmbed{ui: ui, face: face, buttonImage: buttonImage}
}

func LoadButtonImage() *widget.ButtonImage {
	tile := LoadImage(resources.ButtonTileImage)

	images := make([]*image.NineSlice, 4)
	for i := range 4 {
		images[i] = image.NewNineSliceSimple(tile.SubImage(img.Rect(i*16, 0, (i+1)*16, 16)).(*ebiten.Image), 4, 8)
	}
	return &widget.ButtonImage{
		Idle:     images[0],
		Hover:    images[1],
		Pressed:  images[2],
		Disabled: images[3],
	}
}

func LoadProgressBarImage() (*widget.ProgressBarImage, *widget.ProgressBarImage) {
	tile := LoadImage(resources.ProgressBarTileImage)

	w := [3]int{6, 4, 6}
	h := [3]int{7, 2, 7}

	track := image.NewNineSlice(tile.SubImage(img.Rect(0, 0, 16, 16)).(*ebiten.Image), w, h)
	fill := image.NewNineSlice(tile.SubImage(img.Rect(16, 0, 32, 16)).(*ebiten.Image), w, h)

	return &widget.ProgressBarImage{Idle: track}, &widget.ProgressBarImage{Idle: fill}
}

func LoadFont() (font.Face, error) {
	ttfFont, err := truetype.Parse(resources.Font)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
