package title

import (
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

	rootContainer.AddChild(widget.NewButton(
		u.BasicButtonOpts("Play", func(args *widget.ButtonClickedEventArgs) {
			scene.Push("game")
		}),
	))

	return u
}
