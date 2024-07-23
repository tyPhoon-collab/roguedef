package resources

import (
	_ "embed"
)

//go:embed images/ui/button.png
var ButtonImage []byte

//go:embed images/gopher.png
var GopherImage []byte // for testing

//go:embed images/player_back.png
var PlayerBackImage []byte

//go:embed images/player_front.png
var PlayerFrontImage []byte

//go:embed images/pendant8x8.png
var PendantImage []byte

//go:embed images/bullet.png
var BulletImage []byte

//go:embed images/ui/button_tile.png
var ButtonTileImage []byte

//go:embed fonts/x12y16pxMaruMonica.ttf
var Font []byte
