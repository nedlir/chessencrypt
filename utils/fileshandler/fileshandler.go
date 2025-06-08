package fileshandler

import (
	"fmt"
	"os"
)

const (
	MaxFileSize     = 1024 * 1024 // 1 MB
	ChunkSize       = 6
	PartExtension   = ".pgn"
	PartNamePattern = "part_%012d" + PartExtension
)

func ReadFile(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat %q: %w", path, err)
	}
	if info.Size() > MaxFileSize {
		return nil, fmt.Errorf("file too large: %d bytes (max %d)", info.Size(), MaxFileSize)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", path, err)
	}
	return data, nil
}

func WriteFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %q: %w", path, err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("write %q: %w", path, err)
	}
	return nil
}

func CreateDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("mkdir %q: %w", path, err)
	}
	return nil
}

func ListDir(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("read dir %q: %w", path, err)
	}
	return entries, nil
}
