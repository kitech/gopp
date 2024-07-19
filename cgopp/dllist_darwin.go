package cgopp

/*
#include <cxxabi.h>
#include <dlfcn.h>
#include <mach-o/dyld_images.h>
#include <mach-o/loader.h>
#include <mach-o/dyld.h>
#include <mach-o/nlist.h>

#include <stdio.h>
void DyldListSymbols(uint32_t idx) {
    struct mach_header_64* mohdr = (struct mach_header_64*)_dyld_get_image_header(idx);
	uintptr_t slide = (uintptr_t)_dyld_get_image_vmaddr_slide(idx);
	printf("ncmd: %d, szncmd: %d, slide: %ld\n", mohdr->ncmds, mohdr->sizeofcmds, slide);

	uint32_t ncmds = mohdr->ncmds;
	struct load_command* lc = (struct load_command*)((void*)mohdr + sizeof(struct mach_header_64));
	Dl_info info;

	while (ncmds--) {
		printf("%d, cmdty: %d=?%d, cmdsz: %d\n", ncmds, lc->cmd,  LC_SYMTAB, lc->cmdsize);
		if (lc->cmd == LC_SYMTAB) {
			struct symtab_command* symtab = (struct symtab_command*)lc;
			struct nlist_64* symtab_list = (struct nlist_64*)((uintptr_t)mohdr+symtab->symoff);
			char* strtab = (char*)((uintptr_t)mohdr + symtab->stroff);
			printf("i%d, %d, %d\n", ncmds, symtab->symoff, symtab->stroff);

			for (uint32_t j = 0; j < symtab->nsyms; j++) {
				struct nlist_64* x = (struct nlist_64*)((uintptr_t)symtab_list + sizeof(struct nlist_64)*j);
				if (x->n_type & N_STAB) {
					continue;
				}
				dladdr((void*)(x->n_value + slide), &info);
				printf("jv: %d, %p, %s, %p\n", j, x, strtab, (void*)(x->n_value + slide));
			}
		}
		lc = (void*)lc + lc->cmdsize;
	}

	// for (uint32_t i = 0; i < mohdr->ncmds; i++) {
	// 	struct load_command* cmd = (struct load_command*) ((void*)mohdr + sizeof(struct mach_header_64) * (i+1));

	// 	printf("%d, cmdty: %d=?%d, cmdsz: %d\n", i, cmd->cmd,  LC_SYMTAB, cmd->cmdsize);
	// }
}

*/
import "C"

// https://developer.apple.com/library/archive/documentation/System/Conceptual/ManPages_iPhoneOS/man3/dyld.3.html
// https://opensource.apple.com/source/xnu/xnu-4570.71.2/EXTERNAL_HEADERS/mach-o/loader.h.auto.html
// https://medium.com/a-42-journey/nm-otool-everything-you-need-to-know-to-build-your-own-7d4fef3d7507

func DyldImageCount() int {
	rvx := C._dyld_image_count()
	return int(rvx)
}
func DyldImageName(idx int) string {
	rvx := C._dyld_get_image_name(cu32(idx))
	return GoString(rvx)
}
func DyldListSymbols(idx int) {
	C.DyldListSymbols(cu32(idx))
}

// this this public use
func DyldImages() (rets []string) {
	cnt := DyldImageCount()
	for i := 0; i < cnt; i++ {
		rv := DyldImageName(i)
		rets = append(rets, rv)
	}
	return
}
