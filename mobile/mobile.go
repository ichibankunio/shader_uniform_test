package mobile

import (
	// "strconv"
	// "strings"

	"github.com/hajimehoshi/ebiten/v2/mobile"


	"github.com/ichibankunio/shader_uniform_test/game"
)


func init() {
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game

	mobile.SetGame(&game.Game{})
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
