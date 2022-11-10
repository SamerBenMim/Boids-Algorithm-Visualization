package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	// screenWidth, screenHeight = 300, 300
	boidCount        = 400
	viewRadius       = 15
	adjustmentRadius = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWlock  = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, boid := range boids {
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+2), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-2), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y-1), green)
		// screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		// screen.Set(int(boid.position.x), int(boid.position.y+2), green)
		// screen.Set(int(boid.position.x), int(boid.position.y-2), green)
		// screen.Set(int(boid.position.x-3), int(boid.position.y), green)
		// screen.Set(int(boid.position.x), int(boid.position.y+3), green)
		// screen.Set(int(boid.position.x), int(boid.position.y-3), green)

	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
