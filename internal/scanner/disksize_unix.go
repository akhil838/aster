//go:build unix

package scanner

import (
	"io/fs"
	"syscall"
)

// statBlockSize is the POSIX-defined unit for st_blocks (always 512 bytes,
// regardless of the filesystem's actual block size).
const statBlockSize = 512

// diskSize returns the actual on-disk usage of a file by inspecting allocated
// blocks. This correctly handles sparse files (e.g. Docker.raw) whose logical
// size can be much larger than their physical footprint.
func diskSize(info fs.FileInfo) int64 {
	if sys, ok := info.Sys().(*syscall.Stat_t); ok {
		return sys.Blocks * statBlockSize
	}
	return info.Size()
}
