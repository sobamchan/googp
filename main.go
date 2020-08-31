package main

import (
	"flag"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	args := flag.Args()

	dc := gg.NewContext(1200, 628)

	backgroundImage, err := gg.LoadImage(args[0])
	if err != nil {
		panic(errors.Wrap(err, "load background image"))
	}
	backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)
	dc.DrawImage(backgroundImage, 0, 0)

	margin := 12.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 127})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	fontPath := "./data/font/MPLUS1p-ExtraBold.ttf"
	if err := dc.LoadFontFace(fontPath, 90); err != nil {
		panic(errors.Wrap(err, "load font"))
	}
	dc.SetColor(color.White)
	s := args[1]
	maxWidth := float64(dc.Width())
	dc.DrawStringWrapped(s, 50, 200, 0, 0, maxWidth/2, 1.5, gg.AlignLeft)

	if err := dc.SavePNG(args[2]); err != nil {
		panic(errors.Wrap(err, "save png"))
	}
}
