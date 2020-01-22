package main

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
func defaultGrid() [8]Cell {
	row := [8]Cell{
		newcell(0, 0),
		newcell(30, 0),
		newcell(60, 0),
		newcell(90, 0),
		newcell(120, 0),
		newcell(150, 0),
		newcell(180, 0),
		newcell(210, 0),
	}
	return row
}
