package main

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"gong/Ball"
	"gong/Config"
	"gong/Player"
	"gong/gameState"
	"image/color"
	"math"
)

const (
	screenWidth  = 500
	screenHeight = 300
)

var (
	paddleSpeed = 1
	clr         = color.RGBA{}
	defaultFont = basicfont.Face7x13
	state       = GameState.Init()
	p1          = Player.Player{}
	p2          = Player.Player{}
	ball        = Ball.Init(screenWidth/2, screenHeight/2)
)

func (g *Game) DrawPaddle(screen *ebiten.Image, x, y int, clr color.Color) {
	for i := 0; i < x; i++ {
		for j := y; j < y+30; j++ {
			screen.Set(x, j, clr)
		}
	}
}
func Contains(s []ebiten.Key, e ebiten.Key) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	switch state.CurrentGameWindow {
	case GameState.TwoPlayerGame:
		if !state.EndGame {
			p1.UpdatedThisTick = false
			p2.UpdatedThisTick = false

			if state.FirstGameFrame {
				ball.RandomizeVector()
				state.FirstGameFrame = false
			}

			p1.UprateUpdate()
			p2.UprateUpdate()

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
			deltaxP1 := float64(ball.GetPositionX() - p1.GetPosition().GetX())
			deltayP1 := float64(ball.GetPositionY() - p1.GetPosition().GetY())
			deltaxP2 := float64(ball.GetPositionX() - p2.GetPosition().GetX())
			deltayP2 := float64(ball.GetPositionY() - p2.GetPosition().GetY())

			if math.Abs(deltaxP1) <= 5 && math.Abs(deltayP1) <= 30 {
				ball.PaddleBounce(p1.UpRate, true)
			}

			if math.Abs(deltaxP2) <= 5 && math.Abs(deltayP2) <= 30 {
				ball.PaddleBounce(p2.UpRate, false)
			}

			ball.Update()
			ball.XBounce(screenHeight)
			if ball.GetPositionX() <= 0 {
				p2.Score++
				if p2.Score >= 7 {
					state.EndGame = true
				}
				ball = Ball.Init(screenWidth/2, screenHeight/2)
				ball.RandomizeVector()
			}

			if ball.GetPositionX() >= screenWidth {
				p1.Score++
				if p1.Score >= 7 {
					state.EndGame = true
				}
				ball = Ball.Init(screenWidth/2, screenHeight/2)
				ball.RandomizeVector()
			}

		} else {
			if Contains(g.keys, ebiten.KeySpace) {
				state.CurrentGameWindow = GameState.MainMenu
				state.EndGame = false
				state.FirstGameFrame = true
				p2.Score = 0
				p1.Score = 0
			}

		}

	case GameState.MainMenu:
		if Contains(g.keys, ebiten.Key1) {
			state.CurrentGameWindow = GameState.TwoPlayerGame
		} else if Contains(g.keys, ebiten.Key2) {
			state.FirstGameFrame = true
			state.CurrentGameWindow = GameState.SinglePlayerGame
		} else if Contains(g.keys, ebiten.Key3) {
			return errors.New("Closing Game")
		}

	case GameState.SinglePlayerGame:
		if !state.EndGame {
			p1.UpdatedThisTick = false
			p2.UpdatedThisTick = false
			if state.FirstGameFrame {
				ball.RandomizeVector()
				state.FirstGameFrame = false
			}
			p1.UprateUpdate()
			p2.UprateUpdate()
			// Move player 1 up or down
			if Contains(g.keys, ebiten.KeyArrowUp) && p1.GetPosition().GetY() != 0 {
				p1.PositionUp()
			}
			if Contains(g.keys, ebiten.KeyArrowDown) && p1.GetPosition().GetY() != screenHeight-20 {
				p1.PositionDown()
			}

			deltaxP1 := float64(ball.GetPositionX() - p1.GetPosition().GetX())
			deltayP1 := float64(ball.GetPositionY() - p1.GetPosition().GetY())
			deltaxP2 := float64(ball.GetPositionX() - p2.GetPosition().GetX())
			deltayP2 := float64(ball.GetPositionY() - p2.GetPosition().GetY())

			if math.Abs(deltaxP1) <= 5 && math.Abs(deltayP1) <= 40 {
				ball.PaddleBounce(p1.UpRate, true)
			}

			if math.Abs(deltaxP2) <= 5 && math.Abs(deltayP2) <= 40 {
				ball.PaddleBounce(p2.UpRate, false)
			}

			ball.Update()
			ball.XBounce(screenHeight)
			p2.AutoMove(ball, screenHeight, screenWidth)
			if ball.GetPositionX() <= 0 {
				p2.Score++
				if p2.Score >= 7 {
					state.EndGame = true
				}
				ball = Ball.Init(screenWidth/2, screenHeight/2)
				ball.RandomizeVector()
			}
			if ball.GetPositionX() >= screenWidth {
				p1.Score++
				if p1.Score >= 7 {
					state.EndGame = true
				}
				ball = Ball.Init(screenWidth/2, screenHeight/2)
				ball.RandomizeVector()
			}

		} else {
			if Contains(g.keys, ebiten.KeySpace) {
				state.CurrentGameWindow = GameState.MainMenu
				state.EndGame = false
				state.FirstGameFrame = true
				p2.Score = 0
				p1.Score = 0
			}

		}
	default:
		fmt.Println("Something went very wrong")
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	switch state.CurrentGameWindow {
	case GameState.TwoPlayerGame:
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%v : %v | %v : %v", p1.Name, p1.Score, p2.Name, p2.Score))
		if !state.EndGame {
			g.DrawCircle(screen, ball.GetPositionX(), ball.GetPositionY(), 5, clr)
			g.DrawPaddle(screen, p1.GetPosition().GetX(), p1.GetPosition().GetY(), clr)
			g.DrawPaddle(screen, p2.GetPosition().GetX(), p2.GetPosition().GetY(), clr)
			g.DrawLine(screen, screenWidth/2, 0, clr)
		} else {
			width := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "Press space to return"))
			text.Draw(screen, "Made by Dorian Kalaczynski", defaultFont, 0, screenHeight, clr)
			if p1.Score >= 7 {
				s := fmt.Sprintf("%v won!!", p1.Name)
				text.Draw(screen, s, defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to return", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			} else if p2.Score >= 7 {
				v := fmt.Sprintf("%v won!!", p2.Name)
				text.Draw(screen, v, defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to return", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			} else {
				width = fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "How did we get here ?? "))
				text.Draw(screen, "How did we get here??", defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to return", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			}
		}
	case GameState.MainMenu:
		width := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "GONG"))
		nextWidth := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "Select a Game Mode:"))
		text.Draw(screen, "GONG", defaultFont, (screenWidth/2)-width/2, (screenHeight)*0.1, clr)
		text.Draw(screen, "Select a Game Mode:", defaultFont, (screenWidth/2)-nextWidth/2, (screenHeight)*0.2, clr)
		text.Draw(screen, "1.Two Player", defaultFont, (screenWidth/2)-nextWidth/2, (screenHeight)*0.25, clr)
		text.Draw(screen, "2. Single Player", defaultFont, (screenWidth/2)-nextWidth/2, (screenHeight)*0.3, clr)
		text.Draw(screen, "3. Exit", defaultFont, (screenWidth/2)-nextWidth/2, (screenHeight)*0.35, clr)

	case GameState.SinglePlayerGame:
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%v: %v | Ai : %v", p2.Name, p2.Score, p1.Score))
		if !state.EndGame {
			g.DrawCircle(screen, ball.GetPositionX(), ball.GetPositionY(), 5, clr)
			g.DrawPaddle(screen, p1.GetPosition().GetX(), p1.GetPosition().GetY(), clr)
			g.DrawPaddle(screen, p2.GetPosition().GetX(), p2.GetPosition().GetY(), clr)
			g.DrawLine(screen, screenWidth/2, 0, clr)
		} else {
			width := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "Press space to play again"))
			text.Draw(screen, "Made by Dorian Kalaczynski", defaultFont, 0, screenHeight, clr)
			if p1.Score >= 7 {
				text.Draw(screen, "You Lost!!", defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to return", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			} else if p2.Score >= 7 {
				v := fmt.Sprintf("%v won!!", p2.Name)
				text.Draw(screen, v, defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to return", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			} else {
				width = fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "How did we get here ?? "))
				text.Draw(screen, "How did we get here??", defaultFont, (screenWidth/2)-width/2, screenHeight/2, clr)
				text.Draw(screen, "Press space to play again", defaultFont, (screenWidth/2)-width/2, (screenHeight/2)+30, clr)
			}
		}
	default:
		width := fixed.Int26_6.Ceil(font.MeasureString(defaultFont, "How did we get here??"))
		text.Draw(screen, "How did we get here??", defaultFont, (screenWidth/2)-width/2, (screenHeight / 2), clr)
	}
}

func main() {
	cfg, err := Config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	paddleSpeed = cfg.Game.PaddleSpeed
	p1 = Player.Init(screenWidth*0.9, screenHeight/2, paddleSpeed, cfg.Game.PlayerOneAltName)
	p2 = Player.Init(screenWidth*0.1, screenHeight/2, paddleSpeed, cfg.Game.PlayerTwoAltName)
	clr = cfg.Color

	//Setup Config   ^^^

	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Gong: Pong but Go.")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
