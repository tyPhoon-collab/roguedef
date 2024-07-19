package object

import (
	"image/color"
	"roguedef/domain/upgrade"
	"roguedef/system"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type UI struct {
	ui          *ebitenui.UI
	face        font.Face
	buttonImage *widget.ButtonImage
}

func (u *UI) Register(g *Game, o *system.Object) {}

func (u *UI) Draw(screen *ebiten.Image) {
	u.ui.Draw(screen)
}

func (u *UI) Priority() int {
	return 1000
}

func (u *UI) Update() {
	u.ui.Update()
}

func (u *UI) ShowGameOver() chan int {
	ch := make(chan int)

	var removeFunc widget.RemoveWindowFunc

	windowContent := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	windowContent.AddChild(widget.NewButton(
		u.bOpts("Play Again", func(args *widget.ButtonClickedEventArgs) {
			removeFunc()
			ch <- 0
		})...,
	))
	windowContent.AddChild(widget.NewButton(
		u.bOpts("Quit", func(args *widget.ButtonClickedEventArgs) {
			removeFunc()
			ch <- 1
		})...,
	))

	window := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(windowContent),
		widget.WindowOpts.CloseMode(widget.NONE),
	)
	removeFunc = u.ui.AddWindow(window)

	return ch
}

func (u *UI) ShowUpgradeSelection() chan upgrade.Upgrade {
	ch := make(chan upgrade.Upgrade)

	var removeFunc widget.RemoveWindowFunc

	windowContent := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	for _, v := range upgrade.Values() {
		windowContent.AddChild(widget.NewButton(
			u.bOpts(v.String(), func(args *widget.ButtonClickedEventArgs) {
				removeFunc()
				ch <- v
			})...,
		))
	}

	window := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(windowContent),
		widget.WindowOpts.CloseMode(widget.NONE),
	)

	removeFunc = u.ui.AddWindow(window)

	return ch
}

func (u *UI) bOpts(text string, do func(args *widget.ButtonClickedEventArgs)) []widget.ButtonOpt {
	buttonTextColor := &widget.ButtonTextColor{
		Idle: color.RGBA{255, 255, 255, 255},
	}
	return []widget.ButtonOpt{
		widget.ButtonOpts.Text(text, u.face, buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{Left: 20, Right: 20, Top: 5, Bottom: 5}),
		widget.ButtonOpts.Image(u.buttonImage),
		widget.ButtonOpts.ClickedHandler(do),
	}
}

func NewUI() *UI {
	face, err := loadFont()
	if err != nil {
		panic(err)
	}

	buttonImage, err := loadButtonImage()
	if err != nil {
		panic(err)
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	ui := &ebitenui.UI{Container: rootContainer}

	return &UI{ui: ui, face: face, buttonImage: buttonImage}
}

func loadButtonImage() (*widget.ButtonImage, error) {
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

func loadFont() (font.Face, error) {
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
