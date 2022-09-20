package tinychunk_test

import (
	"fmt"
	. "github.com/BitlyTwiser/tinychunk"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// Simply write files, then delete them.
func fileFunc(fileData []byte) error {
	err := os.WriteFile(fmt.Sprintf("./assets/test/%v.txt", uuid.NewString()), fileData, 0700)

	if err != nil {
		return err
	}

	return nil
}

func openFile(fileName string) []byte {
	file, err := os.ReadFile(fmt.Sprintf("./assets/%v", fileName))

	if err != nil {
		log.Printf("Error opening file for test: %v", err.Error())
		os.Exit(1)
	}

	return file
}

func TestOneMegabyteChunking(t *testing.T) {
	t.Log("chunking 1MB file")
	file := openFile("testFile.txt")

	err := Chunk(file, 1, fileFunc)

	assert.Nil(t, err)
}

func TestTwoMegabyteChunking(t *testing.T) {
	t.Log("chunking 2MB file")
	file := openFile("testFile.txt")

	err := Chunk(file, 2, fileFunc)

	assert.Nil(t, err)
}

func TestThreeMegabyteChunking(t *testing.T) {
	t.Log("chunking 3MB file")
	file := openFile("testFile.txt")

	err := Chunk(file, 3, fileFunc)

	assert.Nil(t, err)
}

func TestFourMegabyteChunking(t *testing.T) {
	t.Log("chunking 4MB file")
	file := openFile("testFile.txt")

	err := Chunk(file, 4, fileFunc)

	assert.Nil(t, err)
}
