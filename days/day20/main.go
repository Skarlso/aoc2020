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
func (img *image) rotate() {
	// reverse the matrix
	for i, j := 0, len(img.pixels)-1; i < j; i, j = i+1, j-1 {
		img.pixels[i], img.pixels[j] = img.pixels[j], img.pixels[i]
	}

	// transpose it
	for i := 0; i < len(img.pixels); i++ {
		for j := 0; j < i; j++ {
			img.pixels[i][j], img.pixels[j][i] = img.pixels[j][i], img.pixels[i][j]
		}
	}
}

// flip flips top->bottom
func (img *image) flip() {
	// two loops, one top -> bottom, two bottom -> top. i, j -> swap rows
	for i, j := 0, len(img.pixels)-1; i < j; i, j = i+1, j-1 {
		img.pixels[i], img.pixels[j] = img.pixels[j], img.pixels[i]
	}
}

// checkSides checks if the current image aligns with the rest of the images next to it
// if there are any.
func (img *image) checkSides(field [][]*image, x, y int) bool { return false }

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
	corners := findCorners(images)
	fmt.Println(corners)
	// fmt.Println(field[0][0], field[len(field)-1][len(field[len(field)-1])], field[0][len(field[len(field)-1])], field[len(field)-1][0])
}

// findCorners finds the corners. They don't have to be in order, they will be multiplied together.
func findCorners(images []*image) [][]*image {
	for i := 0; i < len(images); i++ {
		for j := 0; j < len(images); j++ {

		}
	}
	// find upper-left
	// find upper-right
	// find bottom-left
	// find bottom-right
	return nil
}
