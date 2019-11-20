// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
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
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%02x%02x%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, rgba.R, rgba.G, rgba.B)
		}
	}
	fmt.Println("</svg>")
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

//!-
