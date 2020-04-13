package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif" // register gif decoder
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	format := flag.String("format", "jpeg", "output format, jpeg/png/gif")
	flag.Parse()
	if err := convert(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", *format, err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	}
	return fmt.Errorf("Invalid format %v", format)
}
