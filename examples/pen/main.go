// Package main demonstrates usage.
package main

import (
	"io"
	"log"
	"os"

	"github.com/charmbracelet/x/cellbuf"
	"github.com/olamilekan000/x/ansi"
)

func main() {
	pw := cellbuf.NewPenWriter(os.Stdout)
	defer pw.Close() //nolint:errcheck

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(pw, ansi.Wrap(string(data), 10, "")) //nolint:errcheck,gosec
}
