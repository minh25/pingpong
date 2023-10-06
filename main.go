package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/minh25/pingpong/game"
)

func main() {
	g := game.NewGame()
	g.Init()

	for !rl.WindowShouldClose() {
		g.Draw()
	}

	rl.CloseWindow()
}
