package main

import "fmt"

func main() {
	flood := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(floodFill(flood, 1, 1, 2))
}

type Coordinate struct {
	x, y int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}

	q := []*Coordinate{{x: sr, y: sc}}

	for len(q) > 0 {
		levelSize := len(q)

		for range levelSize {
			curr := q[0]
			q = q[1:]

			bfs(image, oldColor, color, *curr, q)
		}
	}

	return image
}

func bfs(img [][]int, oldColor, targetColor int, curr Coordinate, q []*Coordinate) {
	if curr.x < 0 || curr.x >= len(img) || curr.y < 0 || curr.y >= len(img[0]) {
		return
	}

	if img[curr.x][curr.y] != oldColor {
		return
	}

	img[curr.x][curr.y] = targetColor

	bfs(img, oldColor, targetColor, Coordinate{x: curr.x - 1, y: curr.y}, append(q, &Coordinate{x: curr.x - 1, y: curr.y}))
	bfs(img, oldColor, targetColor, Coordinate{x: curr.x + 1, y: curr.y}, append(q, &Coordinate{x: curr.x + 1, y: curr.y}))
	bfs(img, oldColor, targetColor, Coordinate{x: curr.x, y: curr.y - 1}, append(q, &Coordinate{x: curr.x, y: curr.y - 1}))
	bfs(img, oldColor, targetColor, Coordinate{x: curr.x, y: curr.y + 1}, append(q, &Coordinate{x: curr.x, y: curr.y + 1}))
}
