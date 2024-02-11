package GameState

type GameWindow int

const (
	MainMenu GameWindow = iota
	SinglePlayerGame
	TwoPlayerGame
)

type GameState struct {
	FirstGameFrame    bool
	EndGame           bool
	CurrentGameWindow GameWindow
}

func Init() GameState {
	return GameState{EndGame: false, FirstGameFrame: true, CurrentGameWindow: MainMenu}
}
