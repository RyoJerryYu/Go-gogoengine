package main

import (
	"fmt"

	"github.com/RyoJerryYu/gogoengine/engine"
	"github.com/RyoJerryYu/gogoengine/game"
	"github.com/RyoJerryYu/gogoengine/ui"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	ui := ui.NewUI()
	e := engine.NewEngine(
		engine.WithGame(g),
		engine.WithUI(ui),
	)
	e.Run()
}
