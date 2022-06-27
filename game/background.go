package game

import (
	"github.com/elamre/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Star struct {
	x, y        float64
	size        float64
	opacityMax  float64
	opacityNow  float64
	brightening bool
	speed       float64
}

func (s *Star) Update() {
	if s.brightening {
		s.opacityNow += s.speed * s.opacityMax
		if s.opacityNow > s.opacityMax {
			s.brightening = false
		}
	} else {
		s.opacityNow -= s.speed * s.opacityMax
		if s.opacityNow <= 0 {
			s.brightening = true
		}
	}
}

type StarBackground struct {
	stars []Star
}

func getMaxOpacity() float64 {
	random := rand.Float64()
	if random < 0.6 {
		return random / 2
	} else if random < 0.8 {
		return random
	} else {
		return 0.9
	}
}

func getSize() float64 {
	random := rand.Float64()
	if random < 0.2 {
		return 3
	} else if random < 0.4 {
		return 2
	} else {
		return 1
	}
}

func NewStarBackground(amount, width, height float64) StarBackground {
	bckGround := StarBackground{}
	bckGround.stars = make([]Star, int(amount))

	for i := range bckGround.stars {
		bckGround.stars[i] = Star{
			x:           rand.Float64() * width,
			y:           rand.Float64() * height,
			size:        getSize(),
			brightening: rand.Intn(1) == 1,
			opacityMax:  getMaxOpacity(),
			speed:       rand.Float64() / 100,
		}
		bckGround.stars[i].opacityNow = rand.Float64() * bckGround.stars[i].opacityMax
	}
	return bckGround
}

func (b *StarBackground) Update() {
	for i := range b.stars {
		b.stars[i].Update()
	}
}

func (b StarBackground) Draw(g *ebiten.Image) {
	for _, s := range b.stars {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(s.size, s.size)
		op.GeoM.Translate(s.x, s.y)
		op.ColorM.Scale(1, 1, 1, s.opacityNow)
		g.DrawImage(gameutil.Pixel, op)
	}
}
