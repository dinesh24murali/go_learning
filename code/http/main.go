package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// "logWriter" now can use go "Writer" interface as it has a method
// called "Write" that satisfies (correct arguments and return type)
// the "Writer" interface
type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Since "logWriter" satisfies "Writer" interface, we can pass that to the
	// io.Copy function
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
