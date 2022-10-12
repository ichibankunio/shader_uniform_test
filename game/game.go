package game

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var array = []float32{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

}

//go:embed shader.kage
var shaderByte []byte

var shader *ebiten.Shader

type Game struct {
	counter int
}

func init() {

	var err error
	shader, err = ebiten.NewShader(shaderByte)

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.counter ++

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"Array": array,
	}
	screen.DrawRectShader(screenWidth, screenHeight, shader, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// func main() {
// 	ebiten.SetWindowSize(screenWidth, screenHeight)
// 	ebiten.SetWindowTitle("Shader uniform test")
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}
// }
