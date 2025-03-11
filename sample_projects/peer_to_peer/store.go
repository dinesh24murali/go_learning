package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key)) // [20]byte -> []byte -> [:]byte
	hashStr := hex.EncodeToString(hash[:])
	blocksize := 5
	sliceLength := len(hashStr) / blocksize

	paths := make([]string, sliceLength)

	for i := range sliceLength {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}
	return strings.Join(paths, "/")
}

type PathTransformFunc func(string) string

type StoreOps struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOps
}

func NewStore(opts StoreOps) *Store {
	return &Store{
		StoreOps: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathname := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathname, os.ModePerm); err != nil {
		return err
	}

	filename := "Somefilename"

	pathAndFilename := pathname + "/" + filename

	f, err := os.Create(pathAndFilename)

	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	fmt.Printf("written (%d) bytes to disk", n)

	return nil
}
