package cgopp

/*
#define _GNU_SOURCE
#include <dlfcn.h>
#include <link.h>

struct shlib_image_callback_data {
    int32_t idx;
	int32_t cnt;
	const char** names;
};

int shlib_image_callback(struct dl_phdr_info* info, size_t size, void* data) {
    const char* libpath = info->dlpi_name;
    __auto_type d = (struct shlib_image_callback_data*)data;
    if (d->idx < d->cnt) {
	    d->names[d->idx] = libpath;
	    d->idx += 1;
    }
return 0;
}

*/
import "C"
import (
	"log"
	"runtime"
	"unsafe"
)

// this this public use
func DyldImagesInc(file string) (rets []string) {
	if runtime.GOOS == "android" {
		return androidDyldImagesInc(file)
	} else {
		return nonandroidDyldImagesInc(file)
	}
}
func androidDyldImagesInc(file string) (rets []string) {
	//
	return
}
func nonandroidDyldImagesInc(file string) (rets []string) {
	cbdata := C.struct_shlib_image_callback_data{}
	cbdata.idx = 0
	cbdata.cnt = 128
	cbdata.names = (*charptr)(Mallocgc(128 * int(unsafe.Sizeof(usize(0)))))
	C.dl_iterate_phdr(Go2cfnp(C.shlib_image_callback), (voidptr)(&cbdata))
	log.Println(cbdata)

	slc := *((*[1 << 20]voidptr)(voidptr(cbdata.names)))
	for i := 0; i < int(cbdata.idx); i++ {
		log.Println(i, slc[0])
		libpath := GoString(slc[i])
		rets = append(rets, libpath)
	}
	return
}

// this this public use
func DyldSymbolsInc(file string) (rets []string) {
	panic("not impl")
	return
}
