package title

import (
	"roguedef/resources"
	"roguedef/system"

	"github.com/ebitenui/ebitenui/widget"
)

type UI struct {
	*system.UIEmbed
}

func (u *UI) Register(g *Game, o *system.Object) {}

func NewUI(scene *system.Scene) *UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	u := &UI{system.NewUIEmbed(rootContainer)}

	container := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(16),
		)),
	)

	rootContainer.AddChild(container)

	container.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			})),
		u.BasicTextOpts("Rogue Def"),
	))
	container.AddChild(widget.NewGraphic(
		widget.GraphicOpts.Image(system.LoadImage(resources.PlayerFrontImage)),
	))
	container.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			})),
		u.BasicButtonOpts("Play", func(args *widget.ButtonClickedEventArgs) {
			scene.Push("game")
		}),
	))

	return u
}
