# storer

a simple golang []byte read/write storer

[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE)

---


method ```Truncate() (err error)``` will truncate data stored in storer
method ```WriterString(data string) error``` will write data to override exist data
method ```ReadString() (data string, err error)``` will read data stored in storer
method ```Close() (err error)``` will close storer

not contain parse/format data

fileStorre offer a file base storer
