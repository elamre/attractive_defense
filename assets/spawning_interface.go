package assets

type SpawningInterface interface {
	ShouldSpawn() bool
	WaveDone()
	HasSpawned()
}
