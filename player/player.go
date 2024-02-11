package Player

import (
	"gong/Coordinates"
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
