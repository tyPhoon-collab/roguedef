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

func (u *UI) WaitShowUpgradeSelection() upgrade.Upgrade {
	system.TimeScale = 0
	v := <-u.ShowUpgradeSelection()
	system.TimeScale = 1
	return v
}

func (u *UI) ShowGameOver() chan struct{} {
	ch := make(chan struct{})

	c := u.buildGameOverContainer(ch)
	u.Container().AddChild(c)

	return ch
}

func (u *UI) ShowUpgradeSelection() chan upgrade.Upgrade {
	ch := make(chan upgrade.Upgrade)

	c := u.buildUpgradeSelectionContainer(ch)
	u.Container().AddChild(c)

	return ch
}

func NewUI(scene *system.Scene) *UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
	)
	return &UI{system.NewUIEmbed(rootContainer), scene}
}

func (u *UI) buildGameOverContainer(ch chan struct{}) *widget.Container {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	content := widget.NewContainer(
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

	container.AddChild(content)

	content.AddChild(
		u.newPopupButton("Play Again", func(args *widget.ButtonClickedEventArgs) {
			u.scene.Reload()
			ch <- struct{}{}
			u.Container().RemoveChild(content)
		}),
	)
	content.AddChild(
		u.newPopupButton("Quit", func(args *widget.ButtonClickedEventArgs) {
			u.scene.Pop()
			ch <- struct{}{}
			u.Container().RemoveChild(container)
		}),
	)

	return container
}

func (u *UI) buildUpgradeSelectionContainer(ch chan upgrade.Upgrade) *widget.Container {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	content := widget.NewContainer(
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

	container.AddChild(content)

	for _, v := range upgrade.Values() {
		content.AddChild(
			u.newPopupButton(v.String(), func(args *widget.ButtonClickedEventArgs) {
				ch <- v
				u.Container().RemoveChild(container)
			}),
		)
	}

	return container
}

func (u *UI) newPopupButton(text string, do func(args *widget.ButtonClickedEventArgs)) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			},
		)),
		u.BasicButtonOpts(text, do),
	)
}
