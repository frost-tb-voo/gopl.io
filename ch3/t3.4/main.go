// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
		// lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func svg(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, rgba := corner(i+1, j)
			if !(ax >= 0 && ax <= width && ay >= 0 && ay <= height) {
				continue
			}
			bx, by, _ := corner(i, j)
			if !(bx >= 0 && bx <= width && by >= 0 && by <= height) {
				continue
			}
			cx, cy, _ := corner(i, j+1)
			if !(cx >= 0 && cx <= width && cy >= 0 && cy <= height) {
				continue
			}
			dx, dy, _ := corner(i+1, j+1)
			if !(dx >= 0 && dx <= width && dy >= 0 && dy <= height) {
				continue
			}
			// fmt.Fprintf(os.Stderr, "%x\n", rgba)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%02x%02x%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, rgba.R, rgba.G, rgba.B)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, color.RGBA) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, rgba := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, rgba
}

func f(x, y float64) (float64, color.RGBA) {
	r := math.Hypot(x, y) // distance from (0,0)
	zz := math.Sin(r) / r

	rr := uint8(0xff * math.Sin(r))
	bb := uint8(0xff * (1 - math.Sin(r)))
	cc := color.RGBA{rr, 0x00, bb, 0xff}
	return zz, cc
}
