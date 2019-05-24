// Package storer a simple []byte read/write storer
package storer

// Storer storer interface
type Storer interface {
	// Write truncate and write data to file
	Write(data []byte) error
	// Read read data from storer
	Read() (data []byte, err error)
	// Truncate truncate data
	Truncate() (err error)
	// Close close storer
	Close() (err error)
}
