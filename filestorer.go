package storer

import (
	"io/ioutil"
	"os"
)

// fileStorer file base storer
type fileStorer struct {
	file *os.File
}

// MustNewFileStorer force create file storer
// panic if error occur
func MustNewFileStorer(filePath string) Storer {
	fstorer, err := NewFileStorer(filePath)
	if err != nil {
		panic(err)
	}
	return fstorer
}

// NewFileStorer create file base storer
func NewFileStorer(filePath string) (storer Storer, err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err == nil {
		storer = &fileStorer{
			file: file,
		}
	}
	return
}

// Truncate truncate data stored in storer
func (storer *fileStorer) Truncate() (err error) {
	err = storer.file.Truncate(0)
	return
}

// Write write data replace exist data
func (storer *fileStorer) Write(data []byte) error {
	err := storer.Truncate()
	if err != nil {
		return err
	}
	_, err = storer.file.Write(data)
	if err != nil {
		return err
	}
	// seek file point to start
	_, err = storer.file.Seek(0, 0)
	return err
}

// Read read data
func (storer *fileStorer) Read() (data []byte, err error) {
	d, err := ioutil.ReadAll(storer.file)
	if err != nil {
		return []byte{}, err
	}
	// seek file point to start
	_, err = storer.file.Seek(0, 0)
	return d, err
}

// Close close storer
func (storer *fileStorer) Close() (err error) {
	return storer.file.Close()
}
