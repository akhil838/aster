//go:build !unix

package scanner

import "io/fs"

// diskSize falls back to the logical file size on platforms where
// syscall.Stat_t is not available (e.g. Windows).
func diskSize(info fs.FileInfo) int64 {
	return info.Size()
}
