package buildings

import "github.com/hajimehoshi/ebiten/v2"

type UpgradeAble interface {
	GetUpgradeButton() *ebiten.Image
	UpgradeCost() int
	Upgrade()
}
