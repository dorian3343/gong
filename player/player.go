package Player

import (
	"fmt"
	"gong/Coordinates"
	Ball "gong/ball"
	"math"
)

type Player struct {
	Position        Coordinates.Coordinates
	UpRate          int
	MovementRate    int
	UpdatedThisTick bool
	NoUpdateCount   int
	Score           int
}

func Init(x, y, rate int) Player {
	return Player{
		Coordinates.NewCoordinates(x, y), 0, rate, false, 0, 0,
	}
}

func (p *Player) GetPosition() Coordinates.Coordinates {
	return p.Position
}
func (p *Player) SetPosition(x, y int) {
	p.Position.X = x
	p.Position.Y = y
}

func (p *Player) PositionUp() {
	p.SetPosition(p.GetPosition().GetX(), p.GetPosition().GetY()-p.MovementRate)
	if p.UpRate >= 3 {
		p.UpRate = p.UpRate + 1
	}
	p.UpdatedThisTick = true
}

func (p *Player) PositionDown() {
	p.SetPosition(p.GetPosition().GetX(), p.GetPosition().GetY()+p.MovementRate)
	if p.UpRate <= 3 {
		p.UpRate = p.UpRate - 1
	}
	p.UpdatedThisTick = true
}

func (p *Player) UprateUpdate() {
	if !p.UpdatedThisTick {
		p.NoUpdateCount++
		if p.NoUpdateCount > 12 {
			p.UpRate = 0
			p.NoUpdateCount = 0
		}
	}
}

func (p *Player) AutoMove(b Ball.Ball, height, width int) {
	// the distance between the ball and paddle
	deltaX := float64(b.GetXValue() - p.GetPosition().GetX())
	// the height between the ball and paddle

	deltaY := float64(b.GetYValue() - p.GetPosition().GetY())

	//This calculates in how many frames will the ball reach the paddles x axis
	framesToX := math.Abs(deltaX / float64(b.GetVector().GetX()))
	//This calculates the maximum height if it doesnt bounce before reaching the X axis
	possibleApex := framesToX * float64(b.GetVector().GetY())
	fmt.Println(deltaY)
	if math.Abs(deltaY) > (float64(width) / 3) {
		//go to center
		deltaY = float64(width / 2)
	} else {
		// Move the paddle so it adjusts for the ball
		if deltaY > 5 {
			if p.GetPosition().GetY() <= height-20 || p.GetPosition().GetY() <= 10 {
				p.Position.SetY(p.Position.GetY() + p.MovementRate)

			}
		}
		if deltaY < 5 {
			if p.GetPosition().GetY() >= height-20 || p.GetPosition().GetY() >= 10 {
				p.Position.SetY(p.Position.GetY() - p.MovementRate)

			}
		}

	}

	// Predicition system
	if possibleApex < float64(height)*-1 {
		deltaY = float64(b.GetVector().ReverseY()) * framesToX
		//	fmt.Println("--------------\nProbable top bounce ")
	}
	if possibleApex > float64(height) {
		deltaY = float64(b.GetVector().ReverseY()) * framesToX
		//	fmt.Println("--------------\nProbable top bounce ")
	}

}
