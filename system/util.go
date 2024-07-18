package system

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var TimeScale = 1.0
var DeltaTime = time.Second / time.Duration(ebiten.TPS())

func ScaledDeltaTime() time.Duration {
	return time.Duration(float64(DeltaTime) * TimeScale)
}

func DurationToFrameCount(duration time.Duration) int {
	tps := float64(ebiten.TPS())
	floatFrameCount := duration.Seconds() * tps
	frameCount := int(math.Round(floatFrameCount))
	if frameCount == 0 {
		frameCount = 1
	}
	return frameCount
}
