package main

import (
	"fmt"
	"strings"

	"github.com/kitech/gopp"
)

var ctypes = []string{"int", "int64", "float64", "voidptr", "charptr"}

func genfficallmtx() {
	filename := "../cgopp/littleffi_gen.go"
	sb := strings.Builder{}
	sb.WriteString("// oohmyffi\n")
	sb.WriteString("package cgopp\n\n")
	sb.WriteString("import \"log\"\n")
	sb.WriteString("import \"github.com/ebitengine/purego\"\n")
	defer func() {
		err := gopp.SafeWriteFile(filename, []byte(sb.String()), 0644)
		gopp.ErrPrint(err)
	}()

	sb.WriteString("func litfficallgenimpl[RETY any](tycrc uint64, fnptrx uintptr, args...any) RETY {\n")
	sb.WriteString("var rv RETY\n")
	// sb.WriteString("var lenargs uint64 = uint64(len(args))\n")
	// sb.WriteString("switch lenargs {\n")
	sb.WriteString("switch tycrc {\n")

	for i := 0; i < len(ctypes); i++ {
		// sb.WriteString(fmt.Sprintf("case %d:\n", i+1))
		res1 := Combination(ctypes, i+1)
		for j := 0; j < len(res1); j++ {
			res := Permutation(res1[j])
			gopp.Info(i, j, len(res), res)

			for k := 0; k < len(res); k++ {

				// sb.WriteString(arrtocomment(res[k]))
				// sb.WriteString("\n")
				name := arrtoname(res[k])
				crcval := gopp.Crc64Str(name)
				sb.WriteString(fmt.Sprintf("case %d:\n", crcval))
				sb.WriteString(arrtotypedef(res[k]))
				sb.WriteString("\n")
				sb.WriteString(arrtoregfunc(res[k]))
				sb.WriteString("\n")
				sb.WriteString(arrtocall(res[k]))
				sb.WriteString("\n")
				// sb.WriteString(arrtoprmline(res[k]))
				// sb.WriteString("\n")
				// sb.WriteString(fmt.Sprintf("} // end case %d:\n", i+1))
			}
		}
	}

	sb.WriteString("  default:\n")
	sb.WriteString("  log.Println(\"nocare\", tycrc, len(args), voidptr(fnptrx))\n")
	sb.WriteString("} // end switch tycrc\n")
	sb.WriteString("  return rv\n")
	sb.WriteString("}\n")
	// res := Permutation(ctypes)
	// gopp.Info(len(res), res)

}

func arrtotypedef(a []string) string {
	// s1 := arrtoname(a)
	s2 := arrtoargline(a)

	sx := fmt.Sprintf("var fnv func (%s) RETY", s2)
	return sx
}

func arrtocall(a []string) string {
	// s1 := arrtoname(a)
	s2 := arrtoprmline(a)

	sx := fmt.Sprintf("rv = fnv (%s)", s2)
	return sx
}

func arrtoname(a []string) string {
	s := strings.Join(a, "_")
	// s = fmt.Sprintf("%d", gopp.Crc64Str(s))
	return s
}

func arrtocomment(a []string) string {
	return "// " + strings.Join(a, ", ")
}

func arrtoargline(a []string) string {
	s := strings.Join(a, ", ")
	return s
}
func arrtoprmline(a []string) string {
	var v []string
	for i := 0; i < len(a); i++ {
		s := fmt.Sprintf("args[%d].(%s)", i, a[i])
		v = append(v, s)
	}
	return strings.Join(v, ",")
}
func arrtoregfunc(a []string) string {
	// s := arrtoname(a)
	s := fmt.Sprintf("purego.RegisterFunc(&fnv, fnptrx)\n")
	return s
}

////////
//https://github.com/golang-infrastructure/go-enumerate-algorithm

// Permutation 全排列
func Permutation[T any](slice []T) [][]T {
	return _permutation(slice, 0)
}

func _permutation[T any](slice []T, i int) [][]T {
	// 如果已经没有腾挪的余地了，则认为是OK了
	if i+1 == len(slice) {
		result := make([]T, len(slice))
		for index, value := range slice {
			result[index] = value
		}
		return [][]T{result}
	}
	result := make([][]T, 0)
	for j := i; j < len(slice); j++ {
		slice[j], slice[i] = slice[i], slice[j]
		result = append(result, _permutation(slice, i+1)...)
		slice[j], slice[i] = slice[i], slice[j]
	}
	return result
}

// Combination 组合
func Combination[T any](slice []T, n int) [][]T {
	return _combination(slice, n, 0, make([]T, 0))
}

func _combination[T any](slice []T, n int, stepIndex int, selected []T) [][]T {
	// 选够元素了
	if len(selected) == n {
		result := make([]T, 0)
		result = append(result, selected...)
		return [][]T{result}
	}
	// 到达边界了
	if stepIndex >= len(slice) {
		return [][]T{}
	}
	result := make([][]T, 0)
	// 选中当前元素
	selected = append(selected, slice[stepIndex])
	result = append(result, _combination(slice, n, stepIndex+1, selected)...)
	// 不选中当前元素
	selected = selected[0 : len(selected)-1]
	result = append(result, _combination(slice, n, stepIndex+1, selected)...)
	return result
}
