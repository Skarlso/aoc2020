package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
)

// tile is an tile defined by the ID and the matrix representation of the pixels.
// it can rotate and flip itself and give back a given side.
type tile struct {
	id     int64
	pixels [][]string
}

// rotate always rotates in the right direction
func (img *tile) rotate() {
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
func (img *tile) flip() {
	// two loops, one top -> bottom, two bottom -> top. i, j -> swap rows
	for i, j := 0, len(img.pixels)-1; i < j; i, j = i+1, j-1 {
		img.pixels[i], img.pixels[j] = img.pixels[j], img.pixels[i]
	}
}

func (img *tile) top() string {
	return strings.Join(img.pixels[0], "")
}

func (img *tile) bottom() string {
	return strings.Join(img.pixels[len(img.pixels)-1], "")
}

func (img *tile) left() string {
	var result string
	for i := 0; i < len(img.pixels); i++ {
		result += img.pixels[i][0]
	}
	return result
}

func (img *tile) right() string {
	var result string
	for i := 0; i < len(img.pixels); i++ {
		result += img.pixels[i][len(img.pixels[i])-1]
	}
	return result
}

// find matching side, or has matching side, rotates and flips these images
// until it either finds a side which is matching or it doesn't.
func (img *tile) hasMatchingSideWith(other *tile) bool {
	flipped := false
	rotations := 0
	for {
		if rotations == 4 && !flipped {
			flipped = true
			img.flip()
			rotations = 0
		} else if rotations == 4 && flipped {
			return false
		}

		// rotate other and match...
		oFlipped := false
		oRotations := 0
		for {
			// I have to compare all sides... :( This isn't enough. One of the sides should eventually match with any of the other one's sides.
			// Or does it? I compare the rotated sides to all the other rotated sides don't I.
			if img.bottom() == other.bottom() || img.top() == other.top() || img.left() == other.left() || img.right() == other.right() {
				return true
			}
			if oFlipped && oRotations == 4 {
				break
			} else if !oFlipped && oRotations == 4 {
				other.flip()
				oFlipped = true
				oRotations = 0
			}
			other.rotate()
			oRotations++
		}

		img.rotate()
		rotations++
	}
}

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

	var tiles []*tile
	split := bytes.Split(content, []byte("\n"))
	i := &tile{}
	i.pixels = make([][]string, 0)
	for _, line := range split {
		// new tile is coming up, finish up the current one and be on our way.
		if string(line) == "" {
			tiles = append(tiles, i)
			i = &tile{}
			i.pixels = make([][]string, 0)
			continue
		}

		// Read the ID for a tile.
		if strings.Contains(string(line), "Tile") {
			var d int64
			fmt.Sscanf(string(line), "Tile %d:", &d)
			i.id = d
			continue
		}

		pixels := strings.Split(string(line), "")
		i.pixels = append(i.pixels, pixels)
	}

	// construct all possible tiles
	for _, t := range tiles {
		for i := 0; i < 2; i++ {
			for j := 0; j < 4; j++ {
				current := &tile{
					id:     t.id,
					pixels: make([][]string, len(t.pixels)),
				}
				for k := range t.pixels {
					current.pixels[k] = make([]string, len(t.pixels[k]))
					copy(current.pixels[k], t.pixels[k])
				}
				allTiles = append(allTiles, current)
				t.rotate()
			}
			t.flip()
		}
	}

	//fmt.Println("all tiles:", len(allTiles))
	//for _, t := range allTiles {
	//	fmt.Print("id: ", t.id)
	//	fmt.Println(t.pixels)
	//}

	maxGridSize = int(math.Sqrt(float64(len(allTiles) / 8)))
	//fmt.Println("maxGridSize: ", maxGridSize)
	image = make([][]*tile, maxGridSize, maxGridSize)
	for i := range image {
		image[i] = make([]*tile, maxGridSize, maxGridSize)
	}
	visited := make(map[int64]struct{}, 0)
	constructImage(0, 0, visited)
}

var (
	image       [][]*tile
	allTiles    []*tile
	maxGridSize int
)

func constructImage(row, col int, visited map[int64]struct{}) {
	if row == maxGridSize {
		for _, img := range image {
			for _, j := range img {
				fmt.Print(" ", j.id)
			}
			fmt.Println()
		}
		fmt.Println("mult: ", image[0][0].id*image[len(image)-1][0].id*image[0][len(image[0])-1].id*image[len(image)-1][len(image[len(image)-1])-1].id)
		// There are 8 possible correct solutions depending on how the tiles are rotated and flipped.
		// Each will give a correct solution.
		// Which is fine, it will result in the same thing, we can quit right at the first result.
		return
	}
	for _, t := range allTiles {
		if _, ok := visited[t.id]; ok {
			continue
		}

		if row > 0 && image[row-1][col].bottom() != t.top() {
			continue
		}
		if col > 0 && image[row][col-1].right() != t.left() {
			continue
		}
		image[row][col] = t
		visited[t.id] = struct{}{}
		if col == maxGridSize-1 {
			constructImage(row+1, 0, visited)
		} else {
			constructImage(row, col+1, visited)
		}
		delete(visited, t.id)
	}
}
