package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// Open a file
func Open(filename string) (r io.ReadCloser, err error) {
	var f *os.File
	f, err = os.Open(filename)
	if err != nil {
		return
	}

	r = f
	return
}

// ReadLines from an io.Reader
func ReadLines(r io.Reader) (lines []string, err error) {
	var data []byte
	data, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}

	linesRaw := bytes.Split(data, []byte{'\n'})
	lines = make([]string, len(linesRaw))
	for i, raw := range linesRaw {
		lines[i] = string(raw)
	}
	return
}
