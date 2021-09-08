package main

import "fmt"

func search(image [][]int, sr int, sc int, oldColor int, newColor int) {
	image[sr][sc] = newColor
	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		search(image, sr-1, sc, oldColor, newColor)
	}
	if sr+1 < len(image) && image[sr+1][sc] == oldColor {
		search(image, sr+1, sc, oldColor, newColor)
	}
	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		search(image, sr, sc-1, oldColor, newColor)
	}
	if sc+1 < len(image[sr]) && image[sr][sc+1] == oldColor {
		search(image, sr, sc+1, oldColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor != newColor {
		search(image, sr, sc, oldColor, newColor)
	}
	return image
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(floodFill(image, 1, 1, 2))
}
