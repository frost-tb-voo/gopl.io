package main

import (
	"testing"

	single "../../ch8/t8.5/mandelbrot"
	paralell "../../ch8/t8.5/paralell"
)

func TestSingle(tt *testing.T) {
	single.Mandelbrot()
}

func TestParalell(tt *testing.T) {
	paralell.Mandelbrot(1000)
}
