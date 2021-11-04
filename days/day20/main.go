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

func (img *image) top() string {
	return strings.Join(img.pixels[0], "")
}

func (img *image) bottom() string {
	return strings.Join(img.pixels[len(img.pixels)-1], "")
}

func (img *image) left() string {
	var result string
	for i := 0; i < len(img.pixels); i++ {
		result += img.pixels[i][0]
	}
	return result
}

func (img *image) right() string {
	var result string
	for i := 0; i < len(img.pixels); i++ {
		result += img.pixels[i][len(img.pixels[i])-1]
	}
	return result
}

// find matching side, or has matching side, rotates and flips these images
// until it either finds a side which is matching or it doesn't.
func (img *image) hasMatchingSideWith(other *image) bool {
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
	// This won't work... Do, find tile with only two matching sides. We need four of those.
	topLeft := findTopLeftCorner(images)
	bottomLeft := findBottomLeftCorner(images)
	topRight := findTopRightCorner(images)
	bottomRight := findBottomRightCorner(images)
	fmt.Println("mult: ", topLeft.id*topRight.id*bottomLeft.id*bottomRight.id)
}

func findTilesWithOnlyTwoMatchingSides(images []*image) []*image {
	var result []*image

	for i := 0; i < len(images); i++ {

		current := images[i]
		found := false
		for j := 0; j < len(images); j++ {
			next := images[j]
			if current.id == next.id {
				continue
			}

			if current.bottom() == next.bottom() || current.top() == next.top() || current.left() == next.left() || current.right() == next.right() {
			}
		}

		if found {

		}
	}

	return result
}

// findTopLeftCorner divide and conquer.
func findTopLeftCorner(images []*image) *image {
	var result *image
	// it eventually MUST find one.
	// topLeft... meaning, try finding matches for
	for {
		for i := 0; i < len(images); i++ {
			// once we find a matching side it's locked.
			current := images[i]

			found := false
			rotatedFourTimes := false
			flipped := false
			rotations := 0
			for {
				if rotatedFourTimes && flipped {
					break
				}
				if rotatedFourTimes && !flipped {
					rotatedFourTimes = false
					flipped = true
					rotations = 0
					current.flip()
				}

				// compare

				rotations++
				if rotations == 4 {
					rotatedFourTimes = true
				}
				current.rotate()
				if found {
					break
				}
			}
		}
	}
	return result
}
func findBottomLeftCorner(images []*image) *image {
	var result *image
	// it eventually MUST find one.
	for {

	}
	return result
}
func findTopRightCorner(images []*image) *image {
	var result *image
	// it eventually MUST find one.
	for {

	}
	return result
}
func findBottomRightCorner(images []*image) *image {
	var result *image
	// it eventually MUST find one.
	for {

	}
	return result
}

// // findCorners finds the corners. They don't have to be in order, they will be multiplied together.
// func findCorners(images []*image) []*image {
// 	for len(result) != 4 {
// 		// Select one
// 		// Compare it to the rest -- take care of not comparing it to itself.
// 		// If found a match... don't rotate/flip it anymore.
// 		for i := 0; i < len(images); i++ {
// 			current := images[i]
// 			// pick a side -> top
// 			// start matching...
// 			// found match
// 			// don't change and look for next matching -> right side
// 			// found -> put into result
// 			// not found -> back to start but rotate first.
// 			// if all sides rotated, flip then begin again.
// 			// sounds like a backtracking recursive something.

// 			found := false
// 			rotatedFourTimes := false
// 			flipped := false
// 			rotations := 0

// 			for {
// 				if rotatedFourTimes && flipped {
// 					break
// 				}
// 				if rotatedFourTimes && !flipped {
// 					rotatedFourTimes = false
// 					flipped = true
// 					rotations = 0
// 					current.flip()
// 				}

// 				// compare

// 				rotations++
// 				if rotations == 4 {
// 					rotatedFourTimes = true
// 				}
// 				current.rotate()
// 			}
// 			if found {
// 				result = append(result, current)
// 				// remove from images
// 				images = append(images[:i], images[i+1:]...)
// 			}
// 		}
// 	}
// 	// find upper-left
// 	// find upper-right
// 	// find bottom-left
// 	// find bottom-right
// 	return result
// }
