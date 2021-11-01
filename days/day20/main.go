package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// image is an image defined by the ID and the matrix representation of the pixels.
// it can rotate and flip itself and give back a given side.
type image struct {
	id     int
	pixels [][]string
}

// rotate always rotates in the right direction
func (i *image) rotate() {}

// flip flips top->bottom
func (i *image) flip() {}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// format is, ID, some lines, until an empty line
	var images []*image
	tiles := bytes.Split(content, []byte("\n"))
	i := &image{}
	i.pixels = make([][]string, 0)
	for _, line := range tiles {
		// new image is coming up, finish up the current one and be on our way.
		if string(line) == "" {
			images = append(images, i)
			i = &image{}
			i.pixels = make([][]string, 0)
			continue
		}

		// Read the ID for a tile.
		if strings.Contains(string(line), "Tile") {
			var d int
			fmt.Sscanf(string(line), "Tile %d:", &d)
			i.id = d
			continue
		}
	}
}
