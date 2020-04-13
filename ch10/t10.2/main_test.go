package main_test

import (
	"testing"

	"gopl.io/ch10/t10.2/archivereader"
	_ "gopl.io/ch10/t10.2/archivereader/tar"
	_ "gopl.io/ch10/t10.2/archivereader/zip"
)

func TestTarRead(tt *testing.T) {
	files, err := archivereader.Read("./t10.2.tar")
	if err != nil {
		tt.Fatalf("%v\n", err)
		tt.Fail()
	}
	tt.Logf("%v\n", len(files))
	for _, file := range files {
		tt.Logf("%v\n", file.Name)
	}
}

func TestZipRead(tt *testing.T) {
	files, err := archivereader.Read("./t10.2.zip")
	if err != nil {
		tt.Fatalf("%v\n", err)
		tt.Fail()
	}
	tt.Logf("%v\n", len(files))
	for _, file := range files {
		tt.Logf("%v\n", file.Name)
	}
}
