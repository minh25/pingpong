package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

type Sprite interface {
	// Draw draws the sprite
	Draw()
	// Update updates the sprite
	Update()
	// Destroy destroys the sprite
	Destroy()
}

type sprite struct {
	x     int32
	y     int32
	label string
}

var _ Sprite = (*sprite)(nil)

func (s *sprite) Draw() {
  rl.DrawText("The first instance: " + s.label, s.x, s.y, 20, rl.LightGray)
}

func (s *sprite) Update() {
	panic("implement me")
}

func (s *sprite) Destroy() {
	panic("implement me")
}

// NewSprite creates a new sprite
func NewSprite(x, y int32, label string) Sprite {
	return &sprite{
		x:     x,
		y:     y,
		label: label,
	}
}
