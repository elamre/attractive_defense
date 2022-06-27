package world

import "math"

type Targetable interface {
	GetPixelCoordinates() (x, y int)
}

type staticTarget struct {
	x, y int
}

func (s staticTarget) GetPixelCoordinates() (int, int) {
	return s.x, s.y
}

func GetStaticTarget(x, y int) Targetable {
	return &staticTarget{x: x, y: y}
}

func TargetableDistance(OriginPixelX, OriginPixelY int, t Targetable) float64 {
	tX, tY := t.GetPixelCoordinates()
	// Lots of optimalization can be done here
	dX, dY := float64(tX-OriginPixelX), float64(tY-OriginPixelY)
	return math.Sqrt(math.Pow(dX, 2) + math.Pow(dY, 2))
}
