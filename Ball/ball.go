package Ball

import (
	"gong/Coordinates"
	"math/rand"
	"time"
)

type Ball struct {
	position Coordinates.Coordinates
	vector   Coordinates.Coordinates
}

func Init(x, y int) Ball {
	return Ball{position: Coordinates.Coordinates{X: x, Y: y}, vector: Coordinates.Coordinates{X: 0, Y: 0}}
}

func (b *Ball) SetVectorX(x int) {
	b.GetVector().SetX(x)
}

func (b *Ball) SetVectorY(y int) {
	b.GetVector().SetY(y)
}

func (b *Ball) SetVector(x, y int) {
	b.SetVectorX(x)
	b.SetVectorY(y)
}

func (b *Ball) GetVector() *Coordinates.Coordinates {
	return &b.vector
}

func (b *Ball) ReverseVector() {
	b.SetVector(-b.GetVector().GetX(), -b.GetVector().GetY())
}

func (b *Ball) Update() {
	b.position.X += b.GetVector().GetX()
	b.position.Y += b.GetVector().GetY()
}

func (b *Ball) GetPosition() Coordinates.Coordinates {
	return b.position
}

func (b *Ball) GetPositionX() int {
	return b.position.X
}

func (b *Ball) GetPositionY() int {
	return b.position.Y
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
