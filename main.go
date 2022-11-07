package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 1536, 864
	// screenWidth, screenHeight = 300, 300
	boidCount        = 500
	viewRadius       = 13
	adjustmentRadius = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 2][screenHeight + 2]int
	k       Vector2d
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, boid := range boids {
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x-2), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y+2), green)
		screen.Set(int(boid.position.x), int(boid.position.y-2), green)

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
