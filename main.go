package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/shader_uniform_test/game"
)

// ebitenmobile bind -v -target ios -o ./ios/Mobile.xcframework ./mobile
// ebitenmobile bind -v -target ios -o ../oldmaid3d-ios/ios/Mobile.xcframework ./mobile
func main() {
	ebiten.SetWindowSize(640, 480)

	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
