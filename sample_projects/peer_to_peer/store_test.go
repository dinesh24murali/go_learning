package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathtransformFunc(t *testing.T) {
	key := "Momsbestpicture.png"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
}

func TestStore(t *testing.T) {

	otps := StoreOps{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(otps)

	data := bytes.NewReader([]byte("some big file"))

	if err := s.writeStream("mysepcialpicture", data); err != nil {
		t.Error(err)
	}
}
