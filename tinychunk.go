package tinychunk

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

const (
	ONEMB   int = (1 << 20)
	TWOMB   int = (2 << 20)
	THREEMB int = (3 << 20)
	FOURMB  int = (4 << 20)
)

// chunk is the primary chunker (Tis also the only chunker, this is tinychunk after all)
// The input would be the data you are looking to chunk (a large file perhaps), the desires chunk size (In Megabytes),
// and a function that will be performing the work.
func Chunk(data []byte, mbSize int, chunkAction func([]byte) error) error {
	//Calculate the chunks in 2MB segments, make this variable?
	dataSize := len(data)

	var chunkSize int

	switch mbSize {
	case 1:
		chunkSize = ONEMB
	case 2:
		chunkSize = TWOMB
	case 3:
		chunkSize = THREEMB
	case 4:
		chunkSize = FOURMB
	default:
		fmt.Printf("Size must be 1,2,3, or 4MB. Size of %v is not accounted for.", mbSize)
	}

	// Total chunk allotment
	totalChunks := uint(math.Floor(float64(dataSize) / float64(chunkSize)))

	for i := uint(0); i < totalChunks; i++ {
		// Obtain size of byte partition.
		partitionSize := int(math.Min(float64(chunkSize), float64(dataSize-int(i)*chunkSize)))

		// Making newfound buffer to store data.
    buffer := &bytes.Buffer{}

		// Write data to buffer
		writtenBytes, err := buffer.Write(data[:partitionSize])

		if err != nil {
			return err
		}

		if writtenBytes == 0 {
			return errors.New("No bytes written!")
		}

		// Send forth the data after chunking to function.
		err = chunkAction(buffer.Bytes())

		if err != nil {
			return err
		}
	}

	return nil
}
