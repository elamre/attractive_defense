package game

import "math"

type Targetable interface {
	GetPixelCoordinates() (x, y int)
}

func TargetableDistance(OriginPixelX, OriginPixelY int, t Targetable) float64 {
	tX, tY := t.GetPixelCoordinates()
	// Lots of optimalization can be done here
	dX, dY := float64(tX-OriginPixelX), float64(tY-OriginPixelY)
	return math.Sqrt(math.Pow(dX, 2) + math.Pow(dY, 2))
}
