package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type lightTurretGun struct {
	image *ebiten.Image
	level int
}

func (l *lightTurretGun) CanUpgrade() bool {
	return l.level < 4
}
func (l *lightTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_3)
	case 3:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (l *lightTurretGun) Update(target game.Targetable) {

}
func (l *lightTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
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

func newLightTurretGun() *lightTurretGun {
	return &lightTurretGun{
		image: assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_1),
		level: 0,
	}
}

func NewLightTurret(x, y int) *Turret {
	return NewTurret(x, y, newLightTurretGun(), newDefaultBase())
}
