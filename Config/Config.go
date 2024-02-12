package Config

import (
	"gopkg.in/yaml.v2"
	"image/color"
	"os"
)

type config struct {
	Game struct {
		PaddleSpeed      int     `yaml:"paddle_speed"`
		PlayerOneAltName string  `yaml:"player1_alt_name"`
		PlayerTwoAltName string  `yaml:"player2_alt_name"`
		Color            []uint8 `yaml:"color"`
	} `yaml:"game"`
}

type Config struct {
	config
	Color color.RGBA
}

func Init() (Config, error) {
	f, err := os.Open("conf.yaml")
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	// Convert the array to color.Color
	red := cfg.Game.Color[0]
	green := cfg.Game.Color[1]
	blue := cfg.Game.Color[2]
	alpha := uint8(255) // assuming alpha value is 255
	colorArr := color.RGBA{red, green, blue, alpha}

	cfgFinal := Config{config: cfg, Color: colorArr}

	return cfgFinal, nil
}
