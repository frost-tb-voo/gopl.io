// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

type FloatComplex struct {
	Real *big.Float
	Imag *big.Float
}

// const prec uint = 2
// const prec uint = 4
// const prec uint = 8
// const prec uint = 16
// const prec uint = 32
const prec uint = 64

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		// y := float64(py)/height*(ymax-ymin) + ymin
		y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin).SetPrec(prec)
		for px := 0; px < width; px++ {
			// fmt.Fprintf(os.Stderr, "Pos %d, %d\n", px, py)
			// x := float64(px)/width*(xmax-xmin) + xmin
			x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin).SetPrec(prec)
			// z := complex(x, y)
			z := FloatComplex{x, y}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
		// break
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z FloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	// var v FloatComplex
	v := FloatComplex{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		rr := new(big.Float).Sub(new(big.Float).Mul(v.Real, v.Real), new(big.Float).Mul(v.Imag, v.Imag))
		ii := new(big.Float).Add(new(big.Float).Mul(v.Real, v.Imag), new(big.Float).Mul(v.Imag, v.Real))
		v.Real = new(big.Float).Add(rr, z.Real).SetPrec(prec)
		v.Imag = new(big.Float).Add(ii, z.Imag).SetPrec(prec)
		abs := new(big.Float).Add(new(big.Float).Mul(v.Real, v.Real), new(big.Float).Mul(v.Imag, v.Imag))
		if abs.Cmp(big.NewFloat(4)) >= +1 {
			// return color.Gray{255 - contrast*n}
			// zx, _ := z.Real.Float64()
			// zy, _ := z.Imag.Float64()
			// vx, _ := v.Real.Float64()
			// vy, _ := v.Imag.Float64()
			// fmt.Fprintf(os.Stderr, "%f %f %f %f\n", zx, zy, vx, vy)
			return color.RGBA{contrast * n, 0xff - contrast*n, 0xff - contrast*n, 0xff}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
