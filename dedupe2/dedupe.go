package dedupe2

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
)

type fileSize struct {
	size int64 // File size, per stat syscall
	dev  int32 // Device, per stat syscall on Darwin
}

type filesBySizeAndDev map[fileSize]map[string]bool

func bySizeAndDevice(root string) (filesBySizeAndDev, error) {
	// A map of fileSize to set of paths, as implemented via a bool map.
	sds := make(map[fileSize]map[string]bool)
	err := filepath.Walk(root, func(path string, i os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "cannot walk path %s", path)
		}

		if i.IsDir() {
			return nil
		}

		s := &syscall.Stat_t{}
		if err := syscall.Stat(path, s); err != nil {
			return errors.Wrapf(err, "cannot stat path %s", path)
		}

		sd := fileSize{size: i.Size(), dev: s.Dev}
		if _, ok := sds[sd]; !ok {
			sds[sd] = make(map[string]bool)
		}
		sds[sd][path] = true

		return nil
	})
	return sds, err
}

func hashFile(path string) (uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, errors.Wrap(err, "cannot open file")
	}
	defer f.Close()

	h := fnv.New32()
	if _, err := io.Copy(h, f); err != nil {
		// Despite fulfilling the writer interface, hashes never return errors
		// when writing.
		return 0, errors.Wrap(err, "cannot read file")
	}
	return h.Sum32(), nil
}

func hashDedupe(files filesBySizeAndDev) error {
	hp := make(map[uint32]string)
	for _, paths := range files {
		if len(paths) < 2 {
			for path := range paths {
				fmt.Printf("%s has a unique size within its device and thus cannot have duplicates\n", path)
			}
			continue
		}

		for path := range paths {
			h, err := hashFile(path)
			if err != nil {
				return errors.Wrap(err, "cannot hash file")
			}

			target, ok := hp[h]
			if !ok {
				fmt.Printf("%s has unseen hash %#x\n", path, h)
				hp[h] = path
				continue
			}

			fmt.Printf("%s has seen hash %#x - deduplicating\n", path, h)
			if err := os.Remove(path); err != nil {
				return errors.Wrapf(err, "cannot remove %s", path)
			}
			if err := os.Link(target, path); err != nil {
				return errors.Wrapf(err, "cannot link %s to %s", path, target)
			}
		}
	}

	return nil
}

// Dedupe deduplicates files within a path by calculating their hash and hard
// linking them to another file with the same hash on the same device.
func Dedupe(root string) error {
	files, err := bySizeAndDevice(root)
	if err != nil {
		return errors.Wrap(err, "cannot index files by size and device")
	}

	return errors.Wrap(hashDedupe(files), "cannot deduplicate by hash")
}
