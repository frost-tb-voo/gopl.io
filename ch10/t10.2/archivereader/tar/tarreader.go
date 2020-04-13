package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
	"os"

	"gopl.io/ch10/t10.2/archivereader"
)

func init() {
	archivereader.RegisterFormat("POSIX tar", Match, Read)
}

func Match(filePath string) error {
	tarFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)
	for {
		_, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		return err
	}
	return nil
}

func Read(filePath string) ([]archivereader.FileContent, error) {
	tarFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer tarFile.Close()

	var files []archivereader.FileContent
	// Open and iterate through the files in the archive.
	tr := tar.NewReader(tarFile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Contents of %s:\n", hdr.Name)
		var content bytes.Buffer
		if _, err := content.ReadFrom(tr); err != nil {
			log.Fatal(err)
		}
		// fmt.Println()

		files = append(files, archivereader.FileContent{Name: hdr.Name, Body: content.Bytes()})
	}
	return files, nil
}
