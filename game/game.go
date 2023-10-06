package game

type Game interface {
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
	panic("implement me")
}

func (g *game) Destroy() {
	panic("implement me")
}

// NewGame creates a new game
func NewGame() Game {
	return &game{}
}
