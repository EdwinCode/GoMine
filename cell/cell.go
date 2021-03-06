package cell

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var mineblock *ebiten.Image

type Cell struct {
	square *ebiten.Image
	x      float64
	y      float64
}

func initCellSprite() *ebiten.Image {
	var err error
	mineblock, _, err = ebitenutil.NewImageFromFile("minego.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return mineblock
}

func newcell(x float64, y float64) Cell {
	v := initCellSprite()
	return Cell{v, x, y}
}

func (r *Cell) Draw(screen *ebiten.Image) {
	if r.square == nil {
		// New cell image
		r.square = initCellSprite()
	}

	// Create an empty option struct
	opts := &ebiten.DrawImageOptions{}

	// Add the Scale effect to the option struct.
	opts.GeoM.Scale(.36, .36)

	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(r.x, r.y)

	// Draw the square image to the screen with an empty option
	screen.DrawImage(r.square, opts)
}

func defaultSq() Cell {
	return newcell(0, 0)
}
func DefaultGrid(dx, dy int) [][]Cell {
	grid := make([][]Cell, dy)
	for rows := 0; rows < dy; rows++ {
		grid[rows] = make([]Cell, dx)
		for r := range grid[rows] {
			grid[rows][r] = newcell(float64(30*r), float64(30*rows))
		}
	}
	return grid
}
