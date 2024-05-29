package main

import (
	_ "embed"
	"flag"
	"log"
	"strings"

	"github.com/kitech/gopp"
)

//go:embed codetpl/goppgentpl.go
var goppgentpl []byte

//go:generate goppgen main cgo

// using in go generate, to generate some type alias, little functions

// args "goppgen" <package> [cgo]
func main() {
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
	flag.Parse()
	log.Println(flag.Args())

	pkgname := flag.Arg(0)
	cgook := flag.Arg(1) == "cgo"
	cmtpfx := gopp.IfElse2(cgook, "", "// ")

	sb := &strings.Builder{}
	log.Println("codetpl len", len(goppgentpl))
	lines := strings.Split(string(goppgentpl), "\n")

	cgosec := false
	sbout(sb, "package ", pkgname, "\n\n")
	sb.WriteString("// dont modify " + pkgname + "\n")
	sb.WriteString("\n")

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "// begin cgo" {
			cgosec = true
		} else if line == "// end cgo" {
			cgosec = false
		}
		if cgosec {
			sbout(sb, cmtpfx, line, "\n")
		} else {
			sbout(sb, line, "\n")
		}
	}
	log.Println(sb.Len(), sb.String())
	if sb.Len() == 0 {
		log.Println("Why len 0???")
		return
	}

	codefile := "goppgen.go"
	err := gopp.SafeWriteFile(codefile, []byte(sb.String()), 0755)
	gopp.ErrPrint(err, codefile)
	if err == nil {
		log.Println("Wrote", sb.Len(), codefile)
	}
}
func mainx() {
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
	flag.Parse()
	log.Println(flag.Args())

	pkgname := flag.Arg(0)
	cgook := flag.Arg(1) == "cgo"
	cmtpfx := gopp.IfElse2(cgook, "", "// ")

	sb := &strings.Builder{}
	sb.WriteString("package " + pkgname + "\n")

	sb.WriteString("// dont modify " + pkgname + "\n")
	sb.WriteString("\n")

	// import
	sbout(sb, cmtpfx, "/*"+"\n")
	sbout(sb, cmtpfx, "#include <stdint.h>"+"\n")
	sbout(sb, cmtpfx, "*/"+"\n")
	sbout(sb, cmtpfx, "import \"C\""+"\n")

	sb.WriteString("import \"unsafe\"" + "\n")
	sb.WriteString("\n")

	// main
	sb.WriteString("type i32 = int32" + "\n")
	sb.WriteString("type i64 = int64" + "\n")
	sb.WriteString("type u32 = uint32" + "\n")
	sb.WriteString("type u64 = uint64" + "\n")
	sb.WriteString("type f32 = float32" + "\n")
	sb.WriteString("type f64 = float64" + "\n")

	sb.WriteString("type usize = uintptr" + "\n")
	sb.WriteString("type voidptr = unsafe.Pointer" + "\n")

	sbout(sb, cmtpfx, "type cusize = C.uintptr_t"+"\n")
	sbout(sb, cmtpfx, "type cvoidptr = *C.void"+"\n")

	// funcs
	sb.WriteString(`
	func anyptr2uptr[T any](p *T) usize {
		var pp = usize(vptr(p))
		return pp
		}
		`)
	sb.WriteString("\n")
	if cgook {
		sb.WriteString(`
		func anyptr2uptrc[T any](p *T) cuptr{
			var pp = uintptr(vptr(p))
			return cuptr(pp)
		}		
			`)

	}

	log.Println(sb.String())

	codefile := "goppgen.go"
	err := gopp.SafeWriteFile(codefile, []byte(sb.String()), 0755)
	gopp.ErrPrint(err, codefile)
	if err == nil {
		log.Println("Wrote", sb.Len(), codefile)
	}
}

func sbout(sb *strings.Builder, args ...string) {
	for _, arg := range args {
		sb.WriteString(arg)
	}
}

// func anyptr2uptr[T any](p *T) uintptr {
// 	var pp = uintptr(unsafe.Pointer(p))
// 	return pp
// }
// func anyptr2uptrc[T any](p *T) C.uintptr_t {
// 	var pp = uintptr(unsafe.Pointer(p))
// 	return C.uintptr_t(pp)
// }
