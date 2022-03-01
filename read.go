package main

import (
	"os"
	"io/fs"
	"path/filepath"
)

func readFiles(root string) ([]*file, error) {
	files := []*file{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if x := filepath.Ext(path); x != ".txt" && x != ".info" {
			return nil
		}
		f, err := readFile(path)
		if err != nil {
			return err
		}
		files = append(files, f)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readFile(path string) (*file, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &file{
		Path: path,
		Byte: b,
	}, nil
}
