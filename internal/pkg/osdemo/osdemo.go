package osdemo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
// Function creates file if it does not exist.
func WriteToFile(filePath string, text string) (int, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.WriteString(filePath)
}

// GetFileSize func returns size of the file specified by file descriptor.
// Returned size is in bytes.
func GetFileSize(file *os.File) (int64, error) {
	if file == nil {
		return 0, errors.New("Argument is nil: file")
	}

	fi, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("Failed to access file %s", (err.(*os.PathError)).Path)
	}

	return fi.Size(), nil
}

func demoGetFileSize() {
	log.Println("demoGetFileSize()")

	filePath := "./data-vol/demo/os/dummyfile.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	fileSize, err := GetFileSize(file)
	if err != nil {
		log.Println(err)
	}
	fileSizeStr := strconv.FormatInt(fileSize, 10)
	log.Printf("File %s --> Size: %s (bytes)", filePath, fileSizeStr)

	log.Println("~demoGetFileSize()")
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
	demoGetFileSize()
	demoIsFileASymlink()
	fmt.Printf("\n\n~osdemo.ShowDemo()\n\n")
}
