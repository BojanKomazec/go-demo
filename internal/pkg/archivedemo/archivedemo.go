package archivedemo

// Code inspired by https://stackoverflow.com/a/49057861/404421

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// source - path to directory to be compressed
// destination - path to output zip file
func compressToZip(source, destination string) error {
	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(source, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(source))
		zipFile, err := myZip.Create(relPath)
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
	compressToZip("./data-vol/demo/archive_demo/dir_to_archive", "./data-vol/demo/archive_demo/archive.zip")
}
