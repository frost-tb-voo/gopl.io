package main

import (
	"testing"

	single "./mandelbrot"
	paralell "./paralell"
)

func TestSingle(tt *testing.T) {
	single.Mandelbrot()
}

func TestParalell(tt *testing.T) {
	paralell.Mandelbrot(1000)
}
