package main

import (
	"testing"
	"time"
)

func Test(tt *testing.T) {
	interact(time.Duration(time.Second * 10))
}
