package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter(tt *testing.T) {
	c, count := CountingWriter(os.Stdout)
	c.Write([]byte("hello"))
	tt.Log(*count) // "5", = len("hello")

	c, count = CountingWriter(os.Stdout)
	{
		var name = "Dolly"
		fmt.Fprintf(c, "hello, %s", name)
		tt.Log(*count) // "12", = len("hello, Dolly")
	}
	{
		var name = "Sam"
		fmt.Fprintf(c, "hello, %s", name)
		tt.Log(*count)
	}
	{
		var name = "Mr. Bridges"
		fmt.Fprintf(c, "hello, %s", name)
		tt.Log(*count)
	}

	// Output:
	// 5
	// 12
	// 22
	// 40
}
