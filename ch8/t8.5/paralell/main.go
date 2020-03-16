package paralell

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	png.Encode(os.Stdout, Mandelbrot(10)) // NOTE: ignoring errors
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	// width, height          = 256, 256
	// width, height = 1024, 1024
	width, height = 2048, 2048
)

type syncCount struct {
	sync.Mutex
	count int
}

func Mandelbrot(paralell int) *image.RGBA {
	sema := make(chan struct{}, paralell)
	var wg sync.WaitGroup
	// var count syncCount
	done := make(chan struct{})
	// count.Lock()
	// count.count++
	// count.Unlock()
	wg.Add(1)
	go func() {
		wg.Wait()
		close(done)
	}()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	go func() {
		defer func() {
			// count.Lock()
			// count.count--
			// fmt.Printf("count:%v\n", count.count)
			// count.Unlock()
			// fmt.Printf("sema:%v\n", len(sema))
			wg.Done()
		}()
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			// count.Lock()
			// count.count++
			// count.Unlock()
			wg.Add(1)
			go func(py int, y float64) {
				sema <- struct{}{}
				defer func() {
					<-sema
					// count.Lock()
					// count.count--
					// fmt.Printf("count:%v\n", count.count)
					// count.Unlock()
					// fmt.Printf("sema:%v\n", len(sema))
					wg.Done()
				}()
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					cc := mandelbrot(z)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, cc)
				}
			}(py, y)
		}
	}()
	fmt.Printf("Waiting..\n")
	for range done {
		// wait
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
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
