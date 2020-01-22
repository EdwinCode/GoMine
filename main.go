package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// Display the text though the debug function
	ebitenutil.DebugPrint(screen, "This is GoMine")

	// Default grid
	for _, v := range defaultGrid() {
		v.Draw(screen)
	}

	// Get the x, y position of the cursor from the CursorPosition() function
	x, y := ebiten.CursorPosition()

	// Display the information with "X: xx, Y: xx" format
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n X: %d, Y: %d", x, y))

	return nil
}

func main() {
	// Initialize Ebiten, and loop the update() function
	if err := ebiten.Run(update, 320, 240, 2, "GoMine"); err != nil {
		panic(err)
	}
}
