package gopp

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SafeWriteFile is a drop-in replacement for ioutil.WriteFile;
// but SafeWriteFile writes data to a temporary file first and
// only upon success renames that file to filename.
func SafeWriteFile(filename string, data []byte, perm os.FileMode) error {
	// open temp file
	f, err := ioutil.TempFile(filepath.Dir(filename), "tmp")
	if err != nil {
		return err
	}
	err = f.Chmod(perm)
	if err != nil {
		return err
	}
	tmpname := f.Name()

	// write data to temp file
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	if err != nil {
		return err
	}

	return os.Rename(tmpname, filename)
}

func ReadFile(filename string) (string, error) {
	bcc, err := os.ReadFile(filename)
	return string(bcc), err
}

func ReadFileMust(filename string) string {
	bcc, err := os.ReadFile(filename)
	ErrPrint(err, filename)
	return string(bcc)
}

// Enter works
func PauseAk() {
	var c [1]byte
	n, err := os.Stdin.Read(c[:])
	ErrPrint(err, n)
}
