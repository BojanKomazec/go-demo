package testingdemo

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	aFlag = flag.String("f", "", "a flag")
)

func TestMain(m *testing.M) {
	flag.Parse()

	if *aFlag == "" {
		log.Fatal("shouldn't be empty in main")
	} else {
		log.Printf("aFlag = %s", *aFlag)
	}

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
	fmt.Println("TestA")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}

func TestAdd(t *testing.T) {
	if *aFlag == "" {
		t.Fatal("shouldn't be empty in tests")
	}
}
