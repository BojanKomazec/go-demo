package cryptodemo

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

// CalculateSHA1Impl1 func
// 1st implementation of CalculateSHA1
func CalculateSHA1Impl1(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// CalculateSHA1Impl2 calculates file's sha1 hash.
// 2nd implementation of CalculateSHA1
// To verify the result use (on Ubuntu):
//   sha1sum <file>
func CalculateSHA1Impl2(file *os.File) ([]byte, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// CalculateSHA1 func
// 3rd implementation of CalculateSHA1
func CalculateSHA1(file *os.File) ([]byte, error) {
	return CalculateSHA(sha1.New(), file)
}

// CalculateSHA256Impl1 opens a file, calculates its sha256 hash and then closes it.
// 1st implementation of CalculateSHA256Impl2
// This breaks SRP as function should not care about opening/closing file.
// As io.Copy requires *File, another version of CalculateSHA256Impl2 is provided
// below - the one which accepts *File as an argument.
//
func CalculateSHA256Impl1(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// CalculateSHA256Impl2 calculates file's sha256 hash.
// 2nd implementation of CalculateSHA256Impl2
// To verify the result use (on Ubuntu):
//   sha256sum <file>
func CalculateSHA256Impl2(file *os.File) ([]byte, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// CalculateSHA256 func
// 3rd implementation of CalculateSHA256
func CalculateSHA256(file *os.File) ([]byte, error) {
	return CalculateSHA(sha256.New(), file)
}

// CalculateSHA1Impl2 and CalculateSHA256Impl2 are almost identical.
// The only difference is type of hash that's used. So, let's
// unify these two functions into a single one in order to
// satisfy DRY principle.

// CalculateSHA func calculates file's sh1 or sha256 hash,
// depending on the hasher passed as an argument.
func CalculateSHA(h hash.Hash, file *os.File) ([]byte, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func hashCalculationDemo() {
	filePath := "./data-vol/demo/crypto/hashme.bin"
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	//
	// SHA1
	//

	sha1sum, err := CalculateSHA1Impl2(file)
	if err != nil {
		log.Println(err)
	}

	//
	// Demo various representation of hash value (e.g. Base64 encoding)
	//
	// (1) bd1e8fc6b4b493224f6a2c46c972c4ed867db4bd
	// (2) vR6PxrS0kyJPaixGyXLE7YZ9tL0=
	// (3) vR6PxrS0kyJPaixGyXLE7YZ9tL0=
	// (4) bd1e8fc6b4b493224f6a2c46c972c4ed867db4bd

	// %x formats varue as HEX number with lowercase letters
	sha1hex := fmt.Sprintf("%x", sha1sum)
	log.Printf("File: %s --> sha1: %s\n", filePath, sha1hex)

	sha1Base64URLEncoded := base64.URLEncoding.EncodeToString(sha1sum)
	log.Println("sha1 (base64.URLEncoding.EncodeToString) =", sha1Base64URLEncoded)

	sha1Base64StdEncoded := base64.StdEncoding.EncodeToString(sha1sum)
	log.Println("sha1 (base64.StdEncoding.EncodeToString) =", sha1Base64StdEncoded)

	sha1StrEncoded := hex.EncodeToString(sha1sum)
	log.Println("sha1 (hex.EncodeToString) =", sha1StrEncoded)

	//
	// SHA256
	//

	sha256sum, err := CalculateSHA256Impl2(file)
	if err != nil {
		log.Println(err)
	}

	sha256str := fmt.Sprintf("%x", sha256sum)
	log.Printf("File: %s --> sha256: %s\n", filePath, sha256str)

	// SHA1 (CalculateSHA)

	sha1sum, err = CalculateSHA(sha1.New(), file)
	if err != nil {
		log.Println(err)
	}

	sha1str := fmt.Sprintf("%x", sha1sum)
	log.Printf("File: %s --> sha1: %s\n", filePath, sha1str)

	// SHA256 (CalculateSHA)

	sha256sum, err = CalculateSHA(sha256.New(), file)
	if err != nil {
		log.Println(err)
	}

	sha256str = fmt.Sprintf("%x", sha256sum)
	log.Printf("File: %s --> sha256: %s\n", filePath, sha256str)

	// SHA1 (CalculateSHA1 via CalculateSHA)

	sha1sum, err = CalculateSHA1(file)
	if err != nil {
		log.Println(err)
	}

	sha1str = fmt.Sprintf("%x", sha1sum)
	log.Printf("File: %s --> sha1: %s\n", filePath, sha1str)

	// SHA256 (CalculateSHA256 via CalculateSHA)
	sha256sum, err = CalculateSHA256(file)
	if err != nil {
		log.Println(err)
	}

	sha256str = fmt.Sprintf("%x", sha256sum)
	log.Printf("File: %s --> sha1: %s\n", filePath, sha256str)

}

// ShowDemo func
func ShowDemo() {
	hashCalculationDemo()
}
