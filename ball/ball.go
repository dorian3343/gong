package Ball

import (
	"gong/Coordinates"
	"math/rand"
	"time"
)

type Ball struct {
	Position Coordinates.Coordinates
	Vector   Coordinates.Coordinates
}

func Init(x, y int) Ball {
	return Ball{Position: Coordinates.Coordinates{X: x, Y: y}, Vector: Coordinates.Coordinates{X: 0, Y: 0}}
}

func (b *Ball) SetVector(x, y int) {
	b.Vector.X = x
	b.Vector.Y = y
}

func (b *Ball) GetVector() Coordinates.Coordinates {
	return b.Vector
}

func (b *Ball) ReverseVector() {
	b.Vector.X = -b.Vector.X
	b.Vector.Y = -b.Vector.Y
}

func (b *Ball) Update() {
	b.Position.X += b.Vector.X
	b.Position.Y += b.Vector.Y
}

func (b *Ball) GetPosition() Coordinates.Coordinates {
	return b.Position
}

func (b *Ball) GetXValue() int {
	return b.Position.X
}

func (b *Ball) GetYValue() int {
	return b.Position.Y
}

func (b *Ball) PaddleBounce(uprate int, playerone bool) {
	b.SetVector(b.GetVector().ReverseX(), b.GetVector().GetY())
	if playerone {
		if b.GetVector().GetX() > -20 {
			b.SetVector(b.GetVector().GetX()-1, uprate)
		}
	} else {
		if b.GetVector().GetY() > 20 {
			b.SetVector(b.GetVector().GetX()+1, uprate)
		}
	}
}
func (b *Ball) XBounce(height int) {
	if b.GetPosition().GetY() <= 0 || b.GetPosition().GetY() >= height {
		b.SetVector(b.GetVector().GetX(), b.GetVector().ReverseY())
	}

}
func (b *Ball) RandomizeVector() {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(10 - -10) + -10
	r2 := rand.Intn(10 - -19) + -10
	if r1 == 0 {
		r1 = 5
	}

	if r2 == 0 {
		r2 = 5
	}
	b.SetVector(r1, r2)

}
