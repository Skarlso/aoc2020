package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

// tile is an tile defined by the ID and the matrix representation of the pixels.
// it can rotate and flip itself and give back a given side.
type tile struct {
	id     int64
	pixels [][]string
}

// rotate always rotates in the right direction
func (t *tile) rotate() {
	// reverse the matrix
	for i, j := 0, len(t.pixels)-1; i < j; i, j = i+1, j-1 {
		t.pixels[i], t.pixels[j] = t.pixels[j], t.pixels[i]
	}

	// transpose it
	for i := 0; i < len(t.pixels); i++ {
		for j := 0; j < i; j++ {
			t.pixels[i][j], t.pixels[j][i] = t.pixels[j][i], t.pixels[i][j]
		}
	}
}

// flip flips top->bottom
func (t *tile) flip() {
	// two loops, one top -> bottom, two bottom -> top. i, j -> swap rows
	for i, j := 0, len(t.pixels)-1; i < j; i, j = i+1, j-1 {
		t.pixels[i], t.pixels[j] = t.pixels[j], t.pixels[i]
	}
}

func (t *tile) top() string {
	return strings.Join(t.pixels[0], "")
}

func (t *tile) bottom() string {
	return strings.Join(t.pixels[len(t.pixels)-1], "")
}

func (t *tile) left() string {
	var result string
	for i := 0; i < len(t.pixels); i++ {
		result += t.pixels[i][0]
	}
	return result
}

func (t *tile) right() string {
	var result string
	for i := 0; i < len(t.pixels); i++ {
		result += t.pixels[i][len(t.pixels[i])-1]
	}
	return result
}

// copyPixels gets the pixels out from the tile.
func (t *tile) copyPixels() [][]string {
	pixels := make([][]string, len(t.pixels))
	for k := range t.pixels {
		pixels[k] = make([]string, len(t.pixels[k]))
		copy(pixels[k], t.pixels[k])
	}
	return pixels
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

	maxGridSize = int(math.Sqrt(float64(len(allTiles) / 8)))
	image = make([][]*tile, maxGridSize)
	for i := range image {
		image[i] = make([]*tile, maxGridSize)
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
		sea := constructSea()
		count := findSeaMonsters(sea)
		fmt.Println("Rough waters in the monsters habitat: ", count)
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

var (
	mHead        = "..................#."
	mBody        = "#....##....##....###"
	mTail        = ".#..#..#..#..#..#..."
	monsterHead  = regexp.MustCompile(mHead)
	monsterBody  = regexp.MustCompile(mBody)
	monsterTail  = regexp.MustCompile(mTail)
	monsterCount = 15
)

func findSeaMonsters(sea []string) int {
	found := 0
	count := 0
	for i := 0; i < len(sea); i++ {
		count += strings.Count(sea[i], "#")
		body := sea[i]
		for j := 0; j < len(body); j++ {
			if j+len(mBody) <= len(body) {
				chunk := body[j : j+len(mBody)]
				if monsterBody.MatchString(chunk) {
					if sea[i-1] == "" {
						continue
					}
					head := sea[i-1]
					head = head[j : j+len(mHead)]
					if monsterHead.MatchString(head) {
						if sea[i+1] == "" {
							continue
						}
						tail := sea[i+1]
						tail = tail[j : j+len(mTail)]
						if monsterTail.MatchString(tail) {
							found++
						}
					}
				}
			}
		}

	}

	if found > 0 {
		fmt.Println("found monsters: ", found)
		count -= monsterCount * found
		return count
	}
	return 0
}

// constructSea creates a sea by transposing the tiles of the images
// into a single matrix for better handling.
func constructSea() []string {
	offsetDown := 0
	// image count * the size of one tile
	size := len(image) * len(image[0][0].pixels)
	sea := make([]string, size)
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			pixels := image[i][j].copyPixels()
			// remove top and bottom row
			pixels = pixels[1 : len(pixels)-1]
			if len(sea[j+(offsetDown*len(pixels))]) >= size-(2*len(image)) {
				offsetDown++
			}
			for y := 0; y < len(pixels); y++ {
				joined := strings.Join(pixels[y], "")
				// remove right left border
				sea[y+(offsetDown*len(pixels))] += joined[1 : len(joined)-1]
			}
		}
	}
	return sea
}
