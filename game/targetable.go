package game

type Targetable interface {
	GetPixelCoordinates() (x, y int)
}
