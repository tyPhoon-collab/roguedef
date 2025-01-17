package game

import (
	"fmt"
	"image"
	"roguedef/domain"
	"roguedef/domain/upgrade"
	"roguedef/system"
	"time"

	"github.com/ebitenui/ebitenui/widget"
)

type RemoveFunc = func()

type UI struct {
	*system.UIEmbed
	scene               *system.Scene
	damageTextContainer *widget.Container
	phaseText           *widget.Text
	elapsedTimeText     *widget.Text
	expProgressBar      *widget.ProgressBar
	phaseManager        *PhaseManager
	player              *Player

	elapsedTime        time.Duration
	elapsedTimeUpdater *system.Looper
}

func (u *UI) Register(g *Game, o *system.Object) {
	u.phaseManager = g.ObjectByTag("phase_manager").Data.(*PhaseManager)
	u.player = g.ObjectByTag("player").Data.(*Player)

	u.UpdateExpProgressBarMax()
}

func (u *UI) Update() {
	u.elapsedTimeUpdater.Update()
	u.UIEmbed.Update()

	u.phaseText.Label = fmt.Sprintf("Phase: %d", u.phaseManager.Phase())
	u.expProgressBar.SetCurrent(u.player.Exp())
}

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

func (u *UI) UpdateExpProgressBarMax() {
	u.expProgressBar.Max = domain.ExpToNextLevel(u.player.Level())
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

func (u *UI) AddTextAt(x, y int, text string) RemoveFunc {
	t := widget.NewText(u.BasicTextOpts(text))
	tw := t.GetWidget()
	tw.SetLocation(tw.Rect.Add(image.Pt(x, y)))

	u.damageTextContainer.AddChild(t)
	return func() { u.damageTextContainer.RemoveChild(t) }
}

func NewUI(scene *system.Scene) *UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
	)

	u := &UI{UIEmbed: system.NewUIEmbed(rootContainer), scene: scene}

	u.damageTextContainer = widget.NewContainer()

	statusRootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(10)),
		)),
	)

	statusContainer := widget.NewContainer(
		u.BackgroundImage(),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionStart,
			StretchHorizontal:  true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
		)),
	)
	statusUpperContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	u.phaseText = widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionStart,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
		u.BasicTextOpts(""),
	)
	u.elapsedTimeText = widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionEnd,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
		u.BasicTextOpts(u.formatElapsedTimeText()),
	)

	u.elapsedTimeUpdater = system.NewLooper(time.Second, u.updateElapsedTimeText)

	track, fill := system.LoadProgressBarImage()
	u.expProgressBar = widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
		widget.ProgressBarOpts.Values(0, 20, 20),
		widget.ProgressBarOpts.Images(track, fill),
	)

	statusUpperContainer.AddChild(u.phaseText)
	statusUpperContainer.AddChild(u.elapsedTimeText)

	statusContainer.AddChild(statusUpperContainer)
	statusContainer.AddChild(u.expProgressBar)

	statusRootContainer.AddChild(statusContainer)

	rootContainer.AddChild(u.damageTextContainer)
	rootContainer.AddChild(statusRootContainer)

	return u

}

func (u *UI) buildGameOverContainer(ch chan struct{}) *widget.Container {
	container, content := u.newWindowLikeContainer("Game Over")

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

	container.AddChild(content)

	return container
}

func (u *UI) buildUpgradeSelectionContainer(ch chan upgrade.Upgrade) *widget.Container {
	container, content := u.newWindowLikeContainer("Upgrade")

	for _, v := range upgrade.Values() {
		content.AddChild(
			u.newPopupButton(v.String(), func(args *widget.ButtonClickedEventArgs) {
				ch <- v
				u.Container().RemoveChild(container)
			}),
		)
	}

	container.AddChild(content)

	return container
}

func (u *UI) newWindowLikeContainer(title string) (*widget.Container, *widget.Container) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	content := widget.NewContainer(
		u.BackgroundImage(),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)),
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	content.AddChild(
		widget.NewText(
			u.BasicTextOpts(title),
			widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			)),
		),
	)

	return container, content
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

func (u *UI) updateElapsedTimeText() {
	u.elapsedTime += time.Second
	u.elapsedTimeText.Label = u.formatElapsedTimeText()
}

func (u *UI) formatElapsedTimeText() string {
	return "Time: " + system.FormatDuration(u.elapsedTime)
}
