package zip

import (
	"archive/zip"
	"bytes"
	"log"

	"gopl.io/ch10/t10.2/archivereader"
)

func init() {
	archivereader.RegisterFormat("zip", Match, Read)
}

func Match(filePath string) error {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return err
	}
	defer r.Close()
	return err
}

func Read(filePath string) ([]archivereader.FileContent, error) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer r.Close()

	var files []archivereader.FileContent
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		// fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		var content bytes.Buffer
		_, err = content.ReadFrom(rc)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		// fmt.Println()

		files = append(files, archivereader.FileContent{Name: f.Name, Body: content.Bytes()})
	}
	return files, nil
}
