package sprite

type Sprite interface {
	// Draw draws the sprite
	Draw()
	// Update updates the sprite
	Update()
	// Destroy destroys the sprite
	Destroy()
}

type sprite struct {
	// x is the x position of the sprite
	x int
	// y is the y position of the sprite
	y int
	// texture is the texture of the sprite
	texture string
}

var _ Sprite = (*sprite)(nil)

func (s *sprite) Draw() {
	panic("implement me")
}

func (s *sprite) Update() {
	panic("implement me")
}

func (s *sprite) Destroy() {
	panic("implement me")
}

// NewSprite creates a new sprite
func NewSprite(x, y int, texture string) Sprite {
	return &sprite{
		x:       x,
		y:       y,
		texture: texture,
	}
}
