package Player

import (
	"gong/Coordinates"
	Ball "gong/ball"
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

func (p *Player) AutoMove(b Ball.Ball) {
	// the distance between the ball and paddle
	deltaX := float64(b.GetXValue() - p.GetPosition().GetX())
	// the height between the ball and paddle

	deltaY := float64(b.GetYValue() - p.GetPosition().GetY())
	if deltaY > 30 {
		p.Position.SetY(p.Position.GetY() + p.MovementRate)
	}
	if deltaY < 30 {
		p.Position.SetY(p.Position.GetY() - p.MovementRate)
	}

	//slight anti lock fix
	if deltaX < 3 && deltaY < 3 {
		p.Position.SetY(p.Position.GetY() - p.MovementRate*3)
	}
}
