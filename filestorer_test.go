package storer_test

import (
	"github.com/ikuiki/storer"
	"os"
	"testing"
)

// TestFileStorer test filestorer store/read data
func TestFileStorer(t *testing.T) {
	testStorerFilePath := "FileStorerExample.txt"
	defer os.Remove(testStorerFilePath)
	fStorer := storer.MustNewFileStorer(testStorerFilePath)
	dataA := []byte("aaaaaaa")
	dataB := []byte("bbbbsasdfbb")
	// write data a
	err := fStorer.Write(dataA)
	if err != nil {
		t.Fatalf("storer.Write(dataA) error: %v", err)
	}
	// read data a twice, the data read should be same as data a
	readData, err := fStorer.Read()
	if err != nil {
		t.Fatalf("storer.Read() error: %v", err)
	}
	if string(readData) != string(dataA) {
		t.Fatalf("data readed at file [%s] diff with dataA [%s]", readData, dataA)
	}
	// second time read
	readData, err = fStorer.Read()
	if err != nil {
		t.Fatalf("storer.Read() error: %v", err)
	}
	if string(readData) != string(dataA) {
		t.Fatalf("data readed at file [%s] diff with dataA [%s]", readData, dataA)
	}
	// write data a
	err = fStorer.Write(dataA)
	if err != nil {
		t.Fatalf("storer.Write(dataA) error: %v", err)
	}
	// and then write data b
	err = fStorer.Write(dataB)
	if err != nil {
		t.Fatalf("storer.Write(dataB) error: %v", err)
	}
	// data read shoud equal data b
	readData, err = fStorer.Read()
	if err != nil {
		t.Fatalf("storer.Read() error: %v", err)
	}
	if string(readData) != string(dataB) {
		t.Fatalf("data readed at file [%s] diff with dataB [%s]", readData, dataB)
	}
}
