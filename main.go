package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
