package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Game interface {
	// Init initializes the game
	Init()
	// Start starts the game
	Start()
	// Stop stops the game
	Stop()
	// Pause pauses the game
	Pause()
	// Resume resumes the game
	Resume()
	// Restart restarts the game
	Restart()
	// Update updates the game
	Update()
	// Draw draws the game
	Draw()
	// Destroy destroys the game
	Destroy()
}

type game struct {
	// isRunning is the flag to check if the game is running
	isRunning bool
	// isPaused is the flag to check if the game is paused
	isPaused bool
}

var _ Game = (*game)(nil)

func (g *game) Init() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
}

func (g *game) Start() {
	panic("implement me")
}

func (g *game) Stop() {
	panic("implement me")
}

func (g *game) Pause() {
	panic("implement me")
}

func (g *game) Resume() {
	panic("implement me")
}

func (g *game) Restart() {
	panic("implement me")
}

func (g *game) Update() {
	panic("implement me")
}

func (g *game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.EndDrawing()
}

func (g *game) Destroy() {
	panic("implement me")
}

// NewGame creates a new game
func NewGame() Game {
	return &game{}
}
