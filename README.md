# tinyChunk
Tiny library for splicing byte streams into chunks.
[![Go Report Card](https://goreportcard.com/badge/github.com/BitlyTwiser/tinychunk)](https://goreportcard.com/report/github.com/BitlyTwiser/tinychunk)

## Installation:
```go get github.com/BitlyTwiser/tinychunk```

## Example usage and Testing:
- For simple usage, one can look to the testing within the project to gleen how the chunker can operate:
- Simply craft a function and pass said function as an argument itno the chunker. This, along with the data and the size of the chunks (in MB) will chunk your data and perform the given operations upon the chunked byte segments.

```
package tinychunk_test

import (
	"fmt"
	. "github.com/BitlyTwiser/tinychunk"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
  "github.com/google/uuid"
)

// Simply write files, then delete them.
func fileFunc(fileData []byte) error {
  err := os.WriteFile(fmt.Sprintf("./assets/%v.txt", uuid.NewString()), fileData, 0700)

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
```

- After running the test, the "test" folder will hold the files resulting from the testing operation (result of ```ls -la test | sort -h```):
```
-rwx------ 1 steve steve 1048576 Sep 19 18:17 048929cb-9bed-4344-a4aa-c527569e205d.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 3ea9eeb4-2cb5-4923-9914-225afd1ec57b.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 41109604-7aba-458f-a86f-0b742ef270b9.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 41e14135-3e45-4e8f-8a94-ad0c0a30e43a.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 48b16133-a072-4f2e-b6cd-a9871fe52970.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 7ebe9cd6-adb2-4b1a-b7ae-9ea51dca1c86.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 b24f2b92-8266-44fd-a39e-fce8f5393e47.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 b50d9972-66d3-45db-b822-912a8f1fdfd9.txt                                                                                  
-rwx------ 1 steve steve 1048576 Sep 19 18:17 f32c8c47-c096-42eb-b9db-1a6ed51f7acd.txt                                                                                  
-rwx------ 1 steve steve 2097152 Sep 19 18:17 0431062f-d5d6-49bd-a230-70b28fa29d86.txt                                                                                  
-rwx------ 1 steve steve 2097152 Sep 19 18:17 23d9d7fd-5382-419f-b4ec-0a538684a4da.txt                                                                                  
-rwx------ 1 steve steve 2097152 Sep 19 18:17 70e265df-e474-43bb-87ca-bed52c8ba075.txt                                                                                  
-rwx------ 1 steve steve 2097152 Sep 19 18:17 b7a719dc-9291-446c-920b-303f760f0a23.txt                                                                                  
-rwx------ 1 steve steve 3145728 Sep 19 18:17 65a72369-76e3-446e-b05c-80266170e495.txt                                                                                  
-rwx------ 1 steve steve 3145728 Sep 19 18:17 d92a6e72-ced5-42c6-9d51-40c33f6dc2d7.txt                                                                                  
-rwx------ 1 steve steve 3145728 Sep 19 18:17 f710795d-c0ee-4a89-8bb5-9d9316807e93.txt                                                                                  
-rwx------ 1 steve steve 4194304 Sep 19 18:17 394fb35b-9c62-415e-9376-c664ce5364e7.txt                                                                                  
-rwx------ 1 steve steve 4194304 Sep 19 18:17 831ccf16-db5e-4625-8053-be1c71ceefe6.txt    
```

## Seed Data for files:
```
head -c 10MB </dev/urandom >testFile.txt
```
