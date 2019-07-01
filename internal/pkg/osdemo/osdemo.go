package osdemo

import (
	"fmt"
	"os"
)

// CreateDirIfNotExist func
func CreateDirIfNotExist(dir string) error {
	var err error
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return err
}

// IsSymlink func checks if file at the specified path is a symlink.
// It returns true if it is and false if it is not.
func IsSymlink(fileName string) (bool, error) {
	isSymlink := false
	fi, err := os.Lstat(fileName)
	if err == nil {
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			isSymlink = true
			fmt.Printf("File is a symlink: %s", fileName)
		} else {
			fmt.Printf("File is NOT a symlink: %s", fileName)
		}
	}
	return isSymlink, err
}

func demoWriteTextToFile() error {
	filePath := "./data-vol/demo/os/dummyfile.txt"
	fmt.Println("Writing to file: ", filePath)

	// Create the file
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	bytesWritten, err := f.WriteString(filePath)
	fmt.Printf("Wrote %d bytes\n", bytesWritten)

	// f.Sync()
	return nil
}

// ShowDemo func
func ShowDemo() {
	demoWriteTextToFile()
	demoIsFileASymlink()
}
