package main

import (
	"flag"
	"log"
	"strings"

	"github.com/kitech/gopp"
)

//go:generate goppgen main cgo

// using in go generate, to generate some type alias, little functions

// args "goppgen" "package" "cgo"
func main() {
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
	flag.Parse()
	log.Println(flag.Args())

	pkgname := flag.Arg(0)
	cgook := flag.Arg(1) == "cgo"

	sb := strings.Builder{}
	sb.WriteString("package " + pkgname + "\n")

	sb.WriteString("// dont modify " + pkgname + "\n")
	sb.WriteString("\n")

	// import
	if cgook {
		sb.WriteString("/*" + "\n")
		sb.WriteString("#include <stdint.h>" + "\n")
		sb.WriteString("*/" + "\n")
		sb.WriteString("import \"C\"" + "\n")
	}
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
	sb.WriteString("type vptr = unsafe.Pointer" + "\n")

	if cgook {
		sb.WriteString("type cuptr = C.uintptr_t" + "\n")
		sb.WriteString("type cvptr = *C.void" + "\n")
	}

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

// func anyptr2uptr[T any](p *T) uintptr {
// 	var pp = uintptr(unsafe.Pointer(p))
// 	return pp
// }
// func anyptr2uptrc[T any](p *T) C.uintptr_t {
// 	var pp = uintptr(unsafe.Pointer(p))
// 	return C.uintptr_t(pp)
// }
