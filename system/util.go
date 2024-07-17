package system

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var DeltaTime = time.Second / time.Duration(ebiten.TPS())

func DurationToFrameCount(duration time.Duration) int {
	tps := float64(ebiten.ActualTPS())
	floatFrameCount := duration.Seconds() * tps
	frameCount := int(math.Round(floatFrameCount))
	if frameCount == 0 {
		frameCount = 1
	}
	return frameCount
}
