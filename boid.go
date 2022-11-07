package main

import (
	"log"
	"math"
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

func (b *Boid) calcAcceleration() Vector2d {
	upper, lower := b.position.AddValue(viewRadius), b.position.SubstractValue(viewRadius)
	avgPosition, avgVelocity, separation := Vector2d{0, 0}, Vector2d{0, 0}, Vector2d{0, 0}
	count := 0.0
	rWlock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {

			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {

				if dist := boids[otherBoidId].position.distance(boids[otherBoidId].position); dist < viewRadius {

					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Substract(boids[otherBoidId].position).DivideValue(dist))
				}
			}
		}
	}
	rWlock.RUnlock()
	acc := Vector2d{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}

	if count > 0 {
		avgVelocity = avgVelocity.DivideValue(count)
		avgPosition = avgPosition.DivideValue(count)
		accCohesion := avgPosition.Substract(b.position).MultiplyValue(adjustmentRadius)
		accSeparation := separation.MultiplyValue(adjustmentRadius)
		accAlignement := avgVelocity.Substract(b.velocity).MultiplyValue(adjustmentRadius)
		acc = acc.Add(accAlignement).Add(accCohesion)
		log.Println(accSeparation)
		//.Add(accSeparation)
	}
	return acc
}
func (b *Boid) borderBounce(pos, maxBorder float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorder-viewRadius {
		return 1 / float64(pos-maxBorder)
	}
	return 0
}
func (b *Boid) moveOne() {

	acceleration := b.calcAcceleration()

	rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2d{-b.velocity.x, b.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2d{b.velocity.x, -b.velocity.y}
	}
	rWlock.Unlock()
}

func createBoid(id int) {
	b := Boid{
		position: Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2d{rand.Float64()*2 - 1, rand.Float64()*2 - 1},
		id:       id,
	}
	boids[id] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}
