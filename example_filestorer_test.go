package storer_test

import (
	"fmt"
	"github.com/ikuiki/storer"
	"os"
)

// Example_fileStorer test filestorer store/read data
func Example_fileStorer() {
	testStorerFilePath := "testFileStorer.txt"
	defer os.Remove(testStorerFilePath)
	fStorer := storer.MustNewFileStorer(testStorerFilePath)
	dataA := []byte("test data")
	// write data a
	err := fStorer.Write(dataA)
	if err != nil {
		panic(err)
	}
	// read data a twice, the data read should be same as data a
	readData, err := fStorer.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readData))
	// Output:
	// test data
}
