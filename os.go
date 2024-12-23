package gopp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

const (
	Stdinfd  = 0
	Stdoutfd = 1
	Stderrfd = 2
)

const NLUnix = "\n"
const NLWin = "\r\n"
const NLHttp = NLWin // telnet, ftp, email
const NLHtml = "<br/>"

func PathExist(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}
func IsDir(p string) bool {
	fio, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fio.IsDir()
}
func IsRegular(p string) bool {
	fio, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fio.Mode().IsRegular()
}

// RunCmdOut runs a comand and returns the commands output, or an error
func RunCmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	occ, err := cmd.Output() // Output runs the command and returns its standard output.
	// ErrPrint(err, name, arg)
	if err != nil {
		return "", err
	}
	cmd.Wait()
	return string(occ), nil
}

func RunCmdCout(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	occ, err := cmd.CombinedOutput() // Output runs the command and returns its standard output.
	// ErrPrint(err, name, arg)
	if err != nil {
		return "", err
	}
	cmd.Wait()
	return string(occ), nil
}

// try process cmd like safe:
// "ls -v -h" => ["ls", "-v", "-h"]
// or: ["ls -v", "-h"] => ["ls", "-v", "-h"]
// or: ["ls", "-v", "-h"] => ["ls", "-v", "-h"]
// or: ["ls", "-v", "-h"] => ["ls", "-v", "-h"]
// EO=STDERR and STDOUT
// (combined, stdout, stderr)
func RunCmd(wkdir string, args ...string) ([]string, error) {
	if len(wkdir) > 0 {
		cdir, err := os.Getwd()
		ErrPrint(err)
		err = os.Chdir(wkdir)
		ErrPrint(err, wkdir, args)
		defer os.Chdir(cdir)
		if err != nil {
			return nil, err
		}
	}

	// try resplit
	reargs := []string{}
	for _, c := range args {
		res := safeSplit(c)
		if len(res) != 0 {
			reargs = append(reargs, res...)
		}
	}

	// log.Println("sh:", wkdir, reargs)
	c := exec.Command(args[0], reargs[1:]...)
	out, err := c.CombinedOutput()
	ErrPrint(err, wkdir, reargs, string(out))
	return strings.Split(strings.TrimSpace(string(out)), "\n"), err
}

// 不需要等待执行完成就能输出
func RunCmdSout(outfn func(buf []byte), wkdir string, args ...string) error {
	if outfn == nil {
		outfn = func(buf []byte) { println(string(buf)) }
	}
	if len(wkdir) > 0 {
		cdir, err := os.Getwd()
		ErrPrint(err)
		err = os.Chdir(wkdir)
		ErrPrint(err, wkdir, args)
		defer os.Chdir(cdir)
		if err != nil {
			return err
		}
	}

	// try resplit
	reargs := []string{}
	for _, c := range args {
		res := safeSplit(c)
		if len(res) != 0 {
			reargs = append(reargs, res...)
		}
	}

	log.Printf("Runit: %d %+#v\n", len(reargs), reargs)
	c := exec.Command(reargs[0], reargs[1:]...)
	errfp, err := c.StderrPipe()
	ErrPrint(err)
	outfp, err := c.StdoutPipe()
	ErrPrint(err)

	err = c.Start()
	ErrPrint(err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sc := bufio.NewScanner(errfp)
		for sc.Scan() {
			txt := sc.Text()
			txt = "E: " + txt
			outfn([]byte(txt))
		}
	}()
	go func() {
		defer wg.Done()
		sc := bufio.NewScanner(outfp)
		for sc.Scan() {
			txt := sc.Text()
			txt = "O: " + txt
			outfn([]byte(txt))
		}
	}()

	wg.Wait()

	return err
}

// maybe comment: this code does not work if the quoted string does not have spaces.
func safeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func FileSize(p string) int64 {
	fio, err := os.Stat(p)
	if err != nil {
		return 0
	}
	return fio.Size()
}

// if not exist, then create empty file
func Touch(p string) error {
	if !FileExist(p) {
		return ioutil.WriteFile(p, nil, 0644)
	}
	return nil
}

func IsEnvSet(name string) bool {
	val := strings.ToLower(os.Getenv(name))
	return val == "1" || val == "on" || val == "true"
}

func Gopaths() []string {
	gopath := os.Getenv("GOPATH")
	paths := strings.Split(gopath, ":")
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = runtime.GOROOT()
	}
	if goroot != "" {
		paths = append(paths, goroot)
	}
	return paths
}
