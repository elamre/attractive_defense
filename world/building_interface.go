package world

type BuildingInterface interface {
	Targetable
	InflictDamage(damage int)
	Alive() bool
}
