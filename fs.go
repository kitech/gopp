package gopp

import (
	"os"
	"path/filepath"
	// "sync"
)

// gdu good, not easy use as library

// https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
// 返回，目录个数，文件个数，总文件大小
func DirSize(path string) (int64, int64, int64, error) {
	var size int64
	var dircnt int64
	var filecnt int64
	// var mu sync.Mutex

	// Function to calculate size for a given path
	var calculateSize func(string) error
	calculateSize = func(p string) error {
		fileInfo, err := os.Lstat(p)
		if err != nil {
			return err
		}

		// Skip symbolic links to avoid counting them multiple times
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		if fileInfo.IsDir() {
			dircnt++
			entries, err := os.ReadDir(p)
			if err != nil {
				return err
			}
			for _, entry := range entries {
				if err := calculateSize(filepath.Join(p, entry.Name())); err != nil {
					return err
				}
			}
		} else {
			filecnt++
			// mu.Lock()
			size += fileInfo.Size()
			// mu.Unlock()
		}
		return nil
	}

	// Start calculation from the root path
	if err := calculateSize(path); err != nil {
		return 0, 0, 0, err
	}

	return dircnt, filecnt, size, nil
}
