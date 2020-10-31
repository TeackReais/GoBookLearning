package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	time.Second
}