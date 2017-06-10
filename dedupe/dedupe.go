// Package dedupe finds all duplicate files within a directory. It ensures all
// duplicates are hardlinks to the same inodes, where possible.
package dedupe

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

type FileSize struct {
	Device int32 // Darwin only. This is uint64 on Linux
	Size   int64
}

// sizeFilter returns a map of device and filesize to a slice of paths with that
// size and device. Files that are uniquely sized within their device cannot be
// deduplicated via a hardlink, so it's wasteful to hash them.
func sizeFilter(root string) (map[FileSize][]string, error) {
	pd := make(map[FileSize][]string)
	err := filepath.Walk(root, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if i.IsDir() {
			return nil
		}
		f, err := os.Open(p)
		if err != nil {
			return err
		}
		defer f.Close()

		s := &syscall.Stat_t{}
		if err := syscall.Fstat(int(f.Fd()), s); err != nil {
			return err
		}

		id := FileSize{Device: s.Dev, Size: i.Size()}

		if _, ok := pd[id]; !ok {
			pd[id] = make([]string, 0, 1)
		}
		pd[id] = append(pd[id], p)
		return nil
	})
	return pd, err
}

func hashFile(path string) (uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := fnv.New32()
	if _, err := io.Copy(h, f); err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func hashDedupe(paths []string) error {
	seen := make(map[uint32]string)
	for _, p := range paths {
		hash, err := hashFile(p)
		if err != nil {
			return err
		}
		target, ok := seen[hash]
		if !ok {
			// This is the first file of its hash.
			fmt.Printf("%s has new hash %#x\n", p, hash)
			seen[hash] = p
			continue
		}
		fmt.Printf("%s has existing hash %#x - deduping\n", p, hash)
		// We've seen this file's hash before. It's a duplicate.
		if err := os.Remove(p); err != nil {
			return err
		}
		if err := os.Link(target, p); err != nil {
			return err
		}
	}
	return nil
}

// Dedupe finds all files within a path that have the same FNV hash and ensures
// they are all hardlinks to the same inodes on disk.
func Dedupe(root string) error {
	bySize, err := sizeFilter(root)
	if err != nil {
		return err
	}
	for _, potentialDupes := range bySize {
		// This is the only file of its size on its device. It can't have dupes.
		if len(potentialDupes) < 2 {
			fmt.Printf("%s has a unique size. It does not have duplicates\n", potentialDupes[0])
			continue
		}
		if err := hashDedupe(potentialDupes); err != nil {
			return err
		}
	}
	return nil
}
