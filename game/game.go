package game

import (
	"image/color"
	"sync"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/shader_uniform_test/game/raycast"
	"github.com/ichibankunio/shader_uniform_test/game/vec2"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type Game struct {
	counter  int
	renderer *raycast.Renderer
}

func (g *Game) Init() {
	texSize := 64

	g.renderer = &raycast.Renderer{}
	g.renderer.Init(SCREEN_WIDTH, SCREEN_HEIGHT, texSize)

	red := ebiten.NewImage(texSize, texSize)
	red.Fill(color.RGBA{255, 0, 0, 255})
	blue := ebiten.NewImage(texSize, texSize)
	blue.Fill(color.RGBA{0, 0, 255, 255})
	green := ebiten.NewImage(texSize, texSize)
	green.Fill(color.RGBA{0, 255, 0, 255})
	black := ebiten.NewImage(texSize, texSize)
	black.Fill(color.RGBA{0, 0, 0, 255})

	g.renderer.SetTextures([4]*ebiten.Image{
		g.renderer.NewTextureSheet([]*ebiten.Image{green, red}),
		g.renderer.NewTextureSheet([]*ebiten.Image{blue}),
		g.renderer.NewTextureSheet([]*ebiten.Image{black}),
	})
	g.renderer.SetLevel([][]float32{
		{
			1, 1, 2, 2, 2, 2, 1, 1, 1, 1,
			1, 0, 0, 0, 0, 0, 1, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 1, 1,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		},
		{
			1, 0, 0, 0, 0, 0, 1, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 1, 1, 1, 1, 0, 1, 1,
			1, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		},
	}, 10, 10)
	g.renderer.Wld.NewTopView()
	g.renderer.Cam.SetPos(vec2.New(64*10*3/4, 64*10/2))

}

var once sync.Once

func (g *Game) Update() error {
	once.Do(g.Init)

	g.renderer.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen)
	g.counter++

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
