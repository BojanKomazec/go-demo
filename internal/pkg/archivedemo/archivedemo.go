package archivedemo

// Code inspired by https://stackoverflow.com/a/49057861/404421

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// source - path to directory to be compressed
// destination - path to output zip file
func compressToZip(source, destination string, includeRootDir bool) error {
	fmt.Printf("compressToZip(). source = %s, destination = %s\n", source, destination)

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}

	myZip := zip.NewWriter(destinationFile)

	err = filepath.Walk(source, func(filePath string, info os.FileInfo, err error) error {
		fmt.Println("filePath = ", filePath)

		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		var prefix string
		if includeRootDir {
			prefix = filepath.Dir(source)
		} else {
			prefix = strings.TrimPrefix(source, "./")
		}

		relPath := strings.TrimPrefix(filePath, prefix)
		fmt.Println("relPath = ", relPath)

		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		fh.Name = relPath

		// If we myZip.Create(relPath) is used then files in the archive will have the following stat output:
		// 		Access: 1979-12-31 00:00:00.000000000 +0000
		// 		Modify: 1979-12-31 00:00:00.000000000 +0000
		// (Change time would be correct)
		zipFile, err := myZip.CreateHeader(fh)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	err = myZip.Close()
	if err != nil {
		return err
	}

	return nil
}

// ShowDemo func
func ShowDemo() {
	compressToZip("./data-vol/demo/archive_demo/dir_to_archive", "./data-vol/demo/archive_demo/archive1.zip", true)
	compressToZip("./data-vol/demo/archive_demo/dir_to_archive", "./data-vol/demo/archive_demo/archive2.zip", false)
}
