package game

import (
	"roguedef/domain/upgrade"
	"roguedef/system"

	"github.com/ebitenui/ebitenui/widget"
)

type UI struct {
	*system.UIEmbed
	scene *system.Scene
}

func (u *UI) Register(g *Game, o *system.Object) {}

func (u *UI) WaitShowGameOver() {
	system.TimeScale = 0
	<-u.ShowGameOver()
	system.TimeScale = 1
}

func (u *UI) ShowGameOver() chan struct{} {
	ch := make(chan struct{})

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
		u.BasicButtonOpts("Play Again", func(args *widget.ButtonClickedEventArgs) {
			removeFunc()
			u.scene.Reload()
			ch <- struct{}{}
		}),
	))
	windowContent.AddChild(widget.NewButton(
		u.BasicButtonOpts("Quit", func(args *widget.ButtonClickedEventArgs) {
			removeFunc()
			u.scene.Pop()
			ch <- struct{}{}
		}),
	))

	window := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(windowContent),
		widget.WindowOpts.CloseMode(widget.NONE),
	)
	removeFunc = u.UI().AddWindow(window)

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
			u.BasicButtonOpts(v.String(), func(args *widget.ButtonClickedEventArgs) {
				removeFunc()
				ch <- v
			}),
		))
	}

	window := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(windowContent),
		widget.WindowOpts.CloseMode(widget.NONE),
	)

	removeFunc = u.UI().AddWindow(window)

	return ch
}

func NewUI(scene *system.Scene) *UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	return &UI{system.NewUIEmbed(rootContainer), scene}
}
