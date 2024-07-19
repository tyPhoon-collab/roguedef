package upgrade

type Upgrade int

const (
	UpgradeFrequency Upgrade = iota
	UpgradePower
	UpgradeSpeed
)

var upgradeStrings = map[Upgrade]string{
	UpgradeFrequency: "Frequency",
	UpgradePower:     "Power",
	UpgradeSpeed:     "Speed",
}

var upgradeValues = []Upgrade{
	UpgradeFrequency,
	UpgradePower,
	UpgradeSpeed,
}

func (u Upgrade) String() string {
	return upgradeStrings[u]
}

func Values() []Upgrade {
	return upgradeValues
}
