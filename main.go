package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"gong/Player"
	"gong/ball"
	"image/color"
	"math"
)

const (
	screenWidth  = 500
	screenHeight = 300
)

var (
	defaultFont = basicfont.Face7x13
	firstFrame  = true
	endgame     = false
	p1          = Player.Init(screenWidth*0.9, 130, 5)
	p2          = Player.Init(screenWidth*0.1, 130, 5)
	ball        = Ball.Init(screenWidth/2, screenHeight/2)
)

type Game struct {
	keys []ebiten.Key
}

func Contains(s []ebiten.Key, e ebiten.Key) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if !endgame {
		p1.UpdatedThisTick = false
		p2.UpdatedThisTick = false

		if firstFrame {
			ball.RandomizeVector()
			firstFrame = false
		}
		if !p1.UpdatedThisTick {
			p1.NoUpdateCount++
			if p1.NoUpdateCount > 12 {
				p1.UpRate = 0
				p1.NoUpdateCount = 0
			}
		}

		if !p2.UpdatedThisTick {
			p2.NoUpdateCount++
			if p2.NoUpdateCount > 12 {
				p2.UpRate = 0
				p2.NoUpdateCount = 0
			}
		}

		// Move player 1 up or down
		if Contains(g.keys, ebiten.KeyArrowUp) && p1.GetPosition().GetY() != 0 {
			p1.PositionUp()
		}
		if Contains(g.keys, ebiten.KeyArrowDown) && p1.GetPosition().GetY() != screenHeight-20 {
			p1.PositionDown()
		}

		// Move player 2 up or down
		if Contains(g.keys, ebiten.KeyW) && p2.GetPosition().GetY() != 0 {
			p2.PositionUp()
		}
		if Contains(g.keys, ebiten.KeyS) && p2.GetPosition().GetY() != screenHeight-20 {
			p2.PositionDown()
		}
		deltaxP1 := float64(ball.GetXValue() - p1.GetPosition().GetX())
		deltayP1 := float64(ball.GetYValue() - p1.GetPosition().GetY())
		deltaxP2 := float64(ball.GetXValue() - p2.GetPosition().GetX())
		deltayP2 := float64(ball.GetYValue() - p2.GetPosition().GetY())

		if math.Abs(deltaxP1) <= 5 && math.Abs(deltayP1) <= 30 {
			ball.PaddleBounce(p1.UpRate, true)
		}

		if math.Abs(deltaxP2) <= 5 && math.Abs(deltayP2) <= 30 {
			ball.PaddleBounce(p2.UpRate, false)
		}

		ball.Update()
		if ball.Position.GetY() <= 0 || ball.Position.GetY() >= screenHeight {
			ball.SetVector(ball.GetVector().GetX(), ball.GetVector().ReverseY())
		}

		if ball.Position.GetX() <= 0 {
			p2.Score++
			if p2.Score >= 7 {
				endgame = true
			}
			ball = Ball.Init(screenWidth/2, screenHeight/2)
			ball.RandomizeVector()
		}

		if ball.Position.GetX() >= screenWidth {
			p1.Score++
			if p1.Score >= 7 {
				endgame = true
			}
			ball = Ball.Init(screenWidth/2, screenHeight/2)
			ball.RandomizeVector()
		}

	} else {
		if Contains(g.keys, ebiten.KeySpace) {
			endgame = false
			firstFrame = true
			p2.Score = 0
			p1.Score = 0
		}

	}
	return nil
}

func (g *Game) DrawPaddle(screen *ebiten.Image, x, y int, clr color.Color) {
	for i := 0; i < x; i++ {
		for j := y; j < y+30; j++ {
			screen.Set(x, j, clr)
		}
	}
}

func (g *Game) DrawLine(screen *ebiten.Image, x, y int, clr color.Color) {
	for i := 0; i < x; i++ {
		for j := y; j < screenHeight; j++ {
			screen.Set(x, j, clr)
		}
	}
}

func (g *Game) DrawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)
		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))
		screen.Set(x1, y1, clr)
	}
	for r := radius - 1; r >= 1; r-- {
		g.DrawCircle(screen, x, y, r, clr)
	}
}
func (g *Game) Draw(screen *ebiten.Image) {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player one : %v | Player two : %v", p1.Score, p2.Score))
	if !endgame {
		g.DrawCircle(screen, ball.GetXValue(), ball.GetYValue(), 5, white)
		g.DrawPaddle(screen, p1.GetPosition().GetX(), p1.GetPosition().GetY(), white)
		g.DrawPaddle(screen, p2.GetPosition().GetX(), p2.GetPosition().GetY(), white)
		g.DrawLine(screen, screenWidth/2, 0, white)
	} else {
		width := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "Press space to play again"))
		text.Draw(screen, "Made by Dorian Kalaczynski", defaultFont, 0, screenHeight, white)
		if p1.Score >= 7 {
			text.Draw(screen, "Player one won!!", defaultFont, (screenWidth/2)*0.9-width, screenHeight/2, white)
			text.Draw(screen, "Press space to play again", defaultFont, (screenWidth/2)*0.9-width, (screenHeight/2)+30, white)
		} else if p2.Score >= 7 {
			text.Draw(screen, "Player two won!!", defaultFont, (screenWidth/2)*0.9-width, screenHeight/2, white)
			text.Draw(screen, "Press space to play again", defaultFont, (screenWidth/2)*0.9-width, (screenHeight/2)+30, white)
		} else {
			width = fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "How did we get here ?? "))
			text.Draw(screen, "How did we get here??", defaultFont, (screenWidth/2)*0.9-width, screenHeight/2, white)
			text.Draw(screen, "Press space to play again", defaultFont, (screenWidth/2)*0.9-width, (screenHeight/2)+30, white)
		}
	}
}

func main() {
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Gong: Pong but Go.")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
