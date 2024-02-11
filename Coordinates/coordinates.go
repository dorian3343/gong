package Coordinates

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) SetX(x int) {
	c.X = x
}

func (c *Coordinates) SetY(y int) {
	c.Y = y
}

func (c Coordinates) GetX() int {
	return c.X
}

func (c Coordinates) GetY() int {
	return c.Y
}

func (c Coordinates) ReverseY() int {
	return -1 * c.Y
}

func (c Coordinates) ReverseX() int {
	return -1 * c.X
}

func NewCoordinates(x, y int) Coordinates {
	return Coordinates{X: x, Y: y}
}
