// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

type RatComplex struct {
	Real *big.Rat
	Imag *big.Rat
}

func _main() {
	base := 10
	precDigit := 32
	prec := uint(256)
	rr := big.NewRat(1, 3)
	fmt.Fprintf(os.Stderr, "%d\n", prec)
	fmt.Fprintf(os.Stderr, "%s\n", rr.FloatString(precDigit))
	{
		rr, _ = new(big.Rat).SetString(rr.FloatString(precDigit))
		fmt.Fprintf(os.Stderr, "%s\n", rr.FloatString(precDigit))
	}
	{
		ff1, bb, _ := big.ParseFloat(rr.FloatString(precDigit), base, prec, big.ToNearestEven)
		ff1 = ff1.SetPrec(prec)
		fmt.Fprintf(os.Stderr, "%f %d\n", ff1, bb)
		ff2, acc := ff1.Float64()
		fmt.Fprintf(os.Stderr, "%f %s\n", ff2, acc.String())
	}
	{
		creator := big.NewFloat(float64(0))
		creator = creator.SetPrec(prec)
		ff1, bb, _ := creator.Parse(rr.FloatString(precDigit), base)
		ff1 = ff1.SetPrec(prec)
		fmt.Fprintf(os.Stderr, "%f %d\n", ff1, bb)
		ff2, acc := ff1.Float64()
		fmt.Fprintf(os.Stderr, "%f %s\n", ff2, acc.String())
	}
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		fmt.Fprintf(os.Stderr, "py %d\n", py)
		// y := float64(py)/height*(ymax-ymin) + ymin
		y := new(big.Rat).Add(new(big.Rat).Mul(big.NewRat(int64(py), height), new(big.Rat).Sub(big.NewRat(int64(ymax), 1), big.NewRat(int64(ymin), 1))), big.NewRat(int64(ymin), 1))
		for px := 0; px < width; px++ {
			// fmt.Fprintf(os.Stderr, "Pos %d, %d\n", px, py)
			// x := float64(px)/width*(xmax-xmin) + xmin
			x := new(big.Rat).Add(new(big.Rat).Mul(big.NewRat(int64(px), width), new(big.Rat).Sub(big.NewRat(int64(xmax), 1), big.NewRat(int64(xmin), 1))), big.NewRat(int64(xmin), 1))
			// z := complex(x, y)
			z := RatComplex{x, y}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z RatComplex) color.Color {
	const iterations = 200
	const contrast = 15
	// const precDigit = 2
	// const precDigit = 4
	// const precDigit = 8
	// const precDigit = 16
	// const precDigit = 32
	const precDigit = 64

	// round
	z.Real, _ = new(big.Rat).SetString(z.Real.FloatString(precDigit))
	z.Imag, _ = new(big.Rat).SetString(z.Imag.FloatString(precDigit))

	// var v RatComplex
	v := RatComplex{big.NewRat(0, 1), big.NewRat(0, 1)}
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		rr := new(big.Rat).Sub(new(big.Rat).Mul(v.Real, v.Real), new(big.Rat).Mul(v.Imag, v.Imag))
		ii := new(big.Rat).Mul(new(big.Rat).Mul(v.Real, v.Imag), big.NewRat(2, 1))
		v.Real = new(big.Rat).Add(rr, z.Real)
		v.Imag = new(big.Rat).Add(ii, z.Imag)
		// round
		v.Real, _ = new(big.Rat).SetString(v.Real.FloatString(precDigit))
		v.Imag, _ = new(big.Rat).SetString(v.Imag.FloatString(precDigit))
		abs := new(big.Rat).Add(new(big.Rat).Mul(v.Real, v.Real), new(big.Rat).Mul(v.Imag, v.Imag))
		if abs.Cmp(big.NewRat(4, 1)) >= +1 {
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
