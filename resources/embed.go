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

//go:embed images/bullet.png
var BulletImage []byte
