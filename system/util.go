package system

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

// round float to int
func Round(f float64) int {
	return int(math.Round(f))
}

func MulIntByFloat(i int, mul float64) int {
	return Round(float64(i) * mul)
}

// ScaleDuration multiplies duration by mul
func ScaleDuration(duration *time.Duration, mul float64) {
	*duration = time.Duration(float64(*duration) * mul)
}

func LoadImage(embedded []byte) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(embedded))
	if err != nil {
		panic(err)
	}
	return img
}

func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	}
	return fmt.Sprintf("%02d", seconds)
}
