package img

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
)

type Draw struct {
	rgba  *image.RGBA
	color color.Color
	out   bytes.Buffer // image output
	src   image.Image  // image source
}

// New struct
func NewDraw(src io.Reader) (d *Draw, err error) {
	i, err := jpeg.Decode(src)
	if err != nil {
		return
	}
	d = &Draw{
		rgba:  image.NewRGBA(i.Bounds()),
		color: color.Black, // default line color
		src:   i,
	}
	return
}

// Set line color
func (d *Draw) SetColor(c color.Color) *Draw {
	d.color = c
	return d
}

// Draw a rectangle
func (d *Draw) DrawRect(r image.Rectangle, thickness int) *Draw {
	draw.Draw(d.rgba, d.rgba.Bounds(), d.src, d.src.Bounds().Min, draw.Over)
	for i := 0; i < thickness; i++ {
		d.rect(r.Min.X-i, r.Min.Y-i, r.Max.X+i, r.Max.Y+i)
	}
	jpeg.Encode(bufio.NewWriter(&d.out), d.rgba, &jpeg.Options{Quality: 100})
	return d
}

// Output bytes of processed image
func (d *Draw) OutputBytes() []byte {
	return d.out.Bytes()
}

// Draws a rectangle
func (d *Draw) rect(x0, y0, x1, y1 int) {
	d.horizontalLine(x0, x1, y0)
	d.horizontalLine(x0, x1, y1)
	d.verticalLine(x0, y0, y1)
	d.verticalLine(x1, y0, y1)
}

func (d *Draw) horizontalLine(x0, x1, y int) {
	for i := x0; i <= x1; i++ {
		d.dot(i, y)
	}
}

func (d *Draw) verticalLine(x, y0, y1 int) {
	for i := y0; i <= y1; i++ {
		d.dot(x, i)
	}
}

// Draw a dot
func (d *Draw) dot(x, y int) {
	d.rgba.Set(x, y, d.color)
}
