package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(4 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.velocity = b.velocity.Add(b.calcAcceleration()).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	next := b.position.Add(b.velocity)
	if next.x > screenWidth || next.x < 0 {
		b.velocity.x = -b.velocity.x
	}
	if next.y > screenHeight || next.y < 0 {
		b.velocity.y = -b.velocity.y
	}
	b.position = b.position.Add(b.velocity)

}

func createBoid(id int) {
	b := Boid{
		position: Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2d{rand.Float64()*2 - 1, rand.Float64()*2 - 1},
		id:       id,
	}
	boids[id] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = id
	go b.start()
}
