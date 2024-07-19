package domain

import "math"

func ModifyStatusByPhase(status *Status, phase int) {
	weight := float64(phase-1)/10 + 1
	status.Hp = int(math.Round(float64(status.Hp) * weight))
}
