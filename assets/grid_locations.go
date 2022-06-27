package assets

var Surroundings3 = []struct{ X, Y int }{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}
var Surroundings5 = []struct{ X, Y int }{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -2, Y: 0},
	{X: 2, Y: 0},
	{X: 0, Y: 2},
	{X: 0, Y: -2},
	{X: 1, Y: 1},
	{X: -1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: -1},
}

var Surroundings3x3 []struct{ X, Y int }

func init() {
	Surroundings3x3 = make([]struct{ X, Y int }, 0)
	for x := -2; x < 3; x++ {
		for y := -2; y < 3; y++ {
			if x == 0 && y == 0 {
				continue
			}
			Surroundings3x3 = append(Surroundings3x3, struct{ X, Y int }{X: x, Y: y})
		}
	}
}
