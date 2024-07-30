package equip

import (
	"roguedef/resources"
	"roguedef/system"

	"github.com/ebitenui/ebitenui/widget"
)

type UI struct {
	*system.UIEmbed
}

func (u *UI) Register(g *system.Game, o *system.Object) {}

func NewUI(scene *system.Scene) *UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	u := &UI{
		UIEmbed: system.NewUIEmbed(rootContainer),
	}

	rootContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			VerticalPosition: widget.AnchorLayoutPositionStart,
		})),
		u.BasicButtonOpts("Back", func(args *widget.ButtonClickedEventArgs) {
			scene.Pop()
		}),
	))
	rootContainer.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.GraphicOpts.Image(system.LoadImage(resources.PlayerFrontImage)),
	))
	rootContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			VerticalPosition:  widget.AnchorLayoutPositionEnd,
			StretchHorizontal: true,
		})),
		u.BasicButtonOpts("OK", func(args *widget.ButtonClickedEventArgs) {
			scene.Push("game")
		}),
	))

	return u
}
