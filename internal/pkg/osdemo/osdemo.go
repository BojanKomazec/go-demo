package osdemo

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

// IsExist func
func IsExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}

// todo: convert this to unit test
func demoIsExist() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Checking if the following directory exist: ", cwd)
	if IsExist(cwd) {
		log.Println("Directory exists")
	} else {
		log.Println("Directory does not exist")
	}

	nonExistingDir := path.Join(cwd, "non_existing_path_segment")

	log.Println("Checking if the following directory exist: ", nonExistingDir)
	if IsExist(nonExistingDir) {
		log.Println("Directory exists")
	} else {
		log.Println("Directory does not exist")
	}
}

// CreateDirIfNotExist func
func CreateDirIfNotExist(dirPath string) error {
	if !IsExist(dirPath) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// CreateSymlink function creates a symbolic link to a specified file.
func CreateSymlink(filePath, symLinkFilePath string) error {
	if _, err := os.Lstat(symLinkFilePath); err == nil {
		if err := os.Remove(symLinkFilePath); err != nil {
			fmt.Printf("Failed to unlink: %+v\n", err)
			return fmt.Errorf("Failed to unlink: %+v", err)
		}
	} else {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("os.Lstat failed. Error: %s", pathErr.Error())
		} else {
			fmt.Printf("Failed to extract PathError from os.Lstat error")
		}

		if os.IsNotExist(err) {
			fmt.Printf("Failed to check symlink: %+v\n", err)
		}
	}

	if err := os.Symlink(filePath, symLinkFilePath); err != nil {
		if linkErr, ok := err.(*os.LinkError); ok {
			return fmt.Errorf("Creating symlink failed. Error: %s", linkErr.Error())
		}
		return err
	}
	return nil
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

func demoFilePathDir() {
	log.Println("filepath.Dir(\".\") = ", filepath.Dir("."))                                     // .
	log.Println("filepath.Dir(\"..\") = ", filepath.Dir(".."))                                   // .
	log.Println("filepath.Dir(\"./dir\") = ", filepath.Dir("./dir"))                             // .
	log.Println("filepath.Dir(\"./dirA/dirB\") = ", filepath.Dir("./dirA/dirB"))                 // dirA
	log.Println("filepath.Dir(\"/dirA/dirB\") = ", filepath.Dir("/dirA/dirB"))                   // /dirA
	log.Println("filepath.Dir(\"/dirA/dirB/file.txt\") = ", filepath.Dir("/dirA/dirB/file.txt")) // /dirA/dirB
	log.Println("filepath.Dir(\"file.txt\") = ", filepath.Dir("file.txt"))                       // .
}

// WriteToFile func writes text to a file. It returns number of bytes written and error (nil in case of no error).
// Function creates file if it does not exist but it does NOT create a directories in file's path.
// It calls WriteString() which under the hood calls unbuffered POSIX function write().
// Although full path to file can be passed to os.Create(), it does NOT create directory path but only a file!
// Create that path with:
//    dirPath := filepath.Dir(filePath)
//    if err := CreateDirIfNotExist(dirPath); err != nil {
// 	      return 0, err
//    }
func WriteToFile(filePath string, text string) (int, error) {
	f, err := os.Create(filePath)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			log.Printf("os.Create failed. Error: %s", pathErr.Error())
		} else {
			log.Printf("Failed to extract PathError from os.Create error")
		}
		log.Println("Failed to create file.")
		return 0, err
	}
	defer f.Close()

	return f.WriteString(text)
}

// WriteToFileBuffered func
func WriteToFileAtPathBuffered(filePath string, text string) (int, error) {
	f, err := os.Create(filePath)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			log.Printf("os.Create failed. Error: %s", pathErr.Error())
		} else {
			log.Printf("Failed to extract PathError from os.Create error")
		}
		log.Println("Failed to create file.")
		return 0, err
	}
	defer f.Close()

	return WriteToFileBuffered(f, text)
}

// WriteToFileBuffered func
// Function does not create/own f so it shall be caller's responsibiity to call f.Close()
func WriteToFileBuffered(f *os.File, text string) (int, error) {
	w := bufio.NewWriter(f)
	n, err := w.WriteString(text)
	if err != nil {
		return n, err
	}

	err = w.Flush()
	return n, err
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

func demoWriteTextToFile(dirPath string, fileName string) error {
	fmt.Println("demoWriteTextToFile()")
	var err error
	filePath := path.Join(dirPath, fileName)
	if err := CreateDirIfNotExist(dirPath); err != nil {
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
	filePath, err := createAbsPath("./data-vol/demo/os/dummyfile2.txt")
	if err != nil {
		return err
	}
	symLinkFilePath := "./data-vol/demo/os/dummyfile2_sl.txt"
	fmt.Printf("Creating symlink %s --> %s\n", symLinkFilePath, filePath)
	err = os.Symlink(filePath, symLinkFilePath)
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

func demoCreateDirectorySymLink(dirPath string, dirSymLinkPath string) {
	fmt.Println("demoCreateDirectorySymLink()")
	if err := CreateSymlink(dirPath, dirSymLinkPath); err != nil {
		fmt.Println(err)
	}
	fmt.Println("~demoCreateDirectorySymLink()")
}

func createAbsPath(relPath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// fmt.Println("Current working directory:", wd)
	return path.Join(wd, relPath), nil
}

// An absolute path specifies the location of a file or directory from the root directory(/).
// Absolute path starts with /.
// Relative path can start either with ., .. or name of the directory or file.
func demoAbs() {
	alreadyAbsPath := "/data-vol/demo/os/dir1/benchmark.txt"
	if filePath, err := filepath.Abs(alreadyAbsPath); err != nil {
		log.Println(err)
	} else {
		log.Println("Absoulute file path =", filePath)
		// output: /data-vol/demo/os/dir1/benchmark.txt
	}

	relPath := "./data-vol/demo/os/dir1/benchmark.txt"
	if filePath, err := filepath.Abs(relPath); err != nil {
		log.Println(err)
	} else {
		log.Println("Absoulute file path =", filePath)
		// output: /home/user/go/src/github.com/BojanKomazec/go-demo/data-vol/demo/os/dir1/benchmark.txt
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nosdemo.ShowDemo()\n\n")

	dirPath, err := createAbsPath("/data-vol/demo/os/dir1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "dummyfile.txt"
	demoWriteTextToFile(dirPath, fileName)
	demoCreateSymlink()
	demoGetFileSize()
	demoIsFileASymlink()

	dirSymLinkPath, err := createAbsPath("/data-vol/demo/os/dir1SymLink")
	if err != nil {
		fmt.Println(err)
		return
	}

	demoCreateDirectorySymLink(dirPath, dirSymLinkPath)
	demoAbs()
	demoFilePathDir()
	demoIsExist()
	fmt.Printf("\n\n~osdemo.ShowDemo()\n\n")
}
