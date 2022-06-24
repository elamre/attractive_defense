package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type heavyTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
}

func (l *heavyTurretGun) GetUpgradeButton() *ebiten.Image {
	return l.upgradeButton
}

func (l *heavyTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 200 * l.level
}

func (l *heavyTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_3)
	case 3:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (l *heavyTurretGun) Update(target game.Targetable) {

}
func (l *heavyTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	baseX, baseY := dst.GeoM.Element(0, 2), dst.GeoM.Element(1, 2)

	op := &ebiten.DrawImageOptions{}
	mouseXFloat := float64(mouseX) - float64(baseX+32)
	mouseYFloat := float64(mouseY) - float64(baseY+32)

	angle := math.Atan2(mouseYFloat, mouseXFloat)

	op.GeoM.Translate(-64/2, -64/2)
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(64/2, 64/2)
	// Place at correct position
	op.GeoM.Translate(float64(baseX), float64(baseY))
	screen.DrawImage(l.image, op)

}

func newHeavyTurretGun() *heavyTurretGun {
	return &heavyTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiHeavyTurretUpgrade),
		level:         1,
	}
}

func NewHeavyTurret(x, y int) *Turret {
	return NewTurret(x, y, newHeavyTurretGun(), newDefaultBase())
}
