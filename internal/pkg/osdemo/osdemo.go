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

// CreateSymlink function creates a symbolic link to a specified file.
func CreateSymlink(filePath, symLinkFilePath string) error {
	if _, err := os.Lstat(symLinkFilePath); err == nil {
		if err := os.Remove(symLinkFilePath); err != nil {
			fmt.Printf("Failed to unlink: %+v\n", err)
			return fmt.Errorf("Failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("Failed to check symlink: %+v\n", err)
	}
	err := os.Symlink(filePath, symLinkFilePath)
	return err
}

// IsSymlink func checks if file at the specified path is a symlink.
// It returns true if it is and false if it is not.
//
// Mask for the type bits. For regular files, none will be set.
// ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
func IsSymlink(fileName string) (bool, error) {
	var fi os.FileInfo
	var err error

	if fi, err = os.Lstat(fileName); err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true, nil
	}

	return false, nil
}

// WriteToFile func writes text to a file.
// Function reurns an error if path to the file does not exist.
func WriteToFile(filePath string, text string) (int, error) {
	bytesWritten := 0
	var err error

	if f, err := os.Create(filePath); err == nil {
		defer f.Close()
		bytesWritten, err = f.WriteString(filePath)
	}

	return bytesWritten, err
}

func demoWriteTextToFile() error {
	fmt.Println("demoWriteTextToFile()")
	var err error
	filePath := "./data-vol/demo/os/dummyfile.txt"
	if err := CreateDirIfNotExist("./data-vol/demo/os/"); err != nil {
		return err
	}
	fmt.Println("Writing to file: ", filePath)
	bytesWritten, err := WriteToFile(filePath, filePath)
	if err == nil {
		fmt.Printf("Wrote %d bytes\n", bytesWritten)
	} else {
		fmt.Println(err)
	}
	fmt.Println("~demoWriteTextToFile()")
	return err
}

func demoCreateSymlink() error {
	fmt.Println("demoCreateSymlink()")
	filePath := "./data-vol/demo/os/dummyfile2.txt"
	symLinkFilePath := "./data-vol/demo/os/dummyfile2_sl.txt"
	fmt.Printf("Creating symlink %s --> %s\n", symLinkFilePath, filePath)
	err := os.Symlink(filePath, symLinkFilePath)
	if err != nil {
		fmt.Printf("Failed to create symlink %s --> %s. Error: %v\n", symLinkFilePath, filePath, err)
		return err
	}
	fmt.Printf("Created symlink %s --> %s. Error: %v\n", symLinkFilePath, filePath, err)
	fmt.Println("~demoCreateSymlink()")
	return err
}

func demoIsFileASymlink() error {
	fmt.Println("demoIsFileASymlink()")
	filePath := "./data-vol/demo/os/dummyfile3.txt"
	symLinkFilePath := "./data-vol/demo/os/dummyfile3_sl.txt"
	var err error
	fmt.Println("Writing to file: ", filePath)
	if _, err := WriteToFile(filePath, filePath); err != nil {
		return err
	}
	if err := CreateSymlink(filePath, symLinkFilePath); err != nil {
		return err
	}

	if isSymLink, err := IsSymlink(symLinkFilePath); err == nil {
		if isSymLink {
			fmt.Println("File is a symlink:", symLinkFilePath)
		} else {
			fmt.Println("File is NOT a symlink:", symLinkFilePath)
		}
	}
	fmt.Println("~demoIsFileASymlink()")
	return err
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nosdemo.ShowDemo()\n\n")
	demoWriteTextToFile()
	demoCreateSymlink()
	demoIsFileASymlink()
	fmt.Printf("\n\n~osdemo.ShowDemo()\n\n")
}
