package gopp

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
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
	// todo
	// disable input buffering
	// exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var c [1]byte
	n, err := os.Stdin.Read(c[:])
	ErrPrint(err, n)
}

// ///
type HtrespTeer struct {
	urd    io.ReadCloser
	teefns []func([]byte) (int, error)
	outrd  io.Reader
}

var _ io.ReadCloser = (*HtrespTeer)(nil)

func HtrespTeerNew(urdx io.ReadCloser, fns ...func([]byte) (int, error)) *HtrespTeer {
	var urd io.ReadCloser
	switch v := urdx.(type) {
	case io.ReadCloser:
		urd = v
	case io.Reader:
		urd = io.NopCloser(v)
	default:
		Warn(reflect.TypeOf(urdx))
	}
	me := &HtrespTeer{urd, fns, nil}
	me.outrd = io.TeeReader(me.urd, me)
	return me
}
func (me *HtrespTeer) Read(b []byte) (int, error) {
	return me.outrd.Read(b)
}
func (me *HtrespTeer) Write(b []byte) (int, error) {
	for _, teefn := range me.teefns {
		if teefn == nil {
			continue
		}
		n, err := teefn(b)
		ErrPrint(err, n)
	}
	return len(b), nil
}
func (me *HtrespTeer) Close() error { return me.urd.Close() }

/////
