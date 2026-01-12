package main

func main() {

}

type Coordinate struct {
	x, y int
}

// BFS - QUEUE

// BIG O - O(mxn) -> pior caso quando tem que colorir tudo
// Space - O(mxn) -> idem

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]

	// ja ta na cor certa
	if oldColor == color {
		return image
	}

	q := []Coordinate{{
		x: sr,
		y: sc,
	}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if image[curr.x][curr.y] == oldColor {
			image[curr.x][curr.y] = color
			// precisamos saber se tem como ir para o caminho
			// if len(image[sr]) > 0
			if curr.x-1 >= 0 && image[curr.x-1][curr.y] == oldColor {
				q = append(q, Coordinate{
					x: curr.x - 1,
					y: curr.y,
				})
			}

			if curr.x+1 < len(image) && image[curr.x+1][curr.y] == oldColor {
				q = append(q, Coordinate{
					x: curr.x + 1,
					y: curr.y,
				})
			}

			if curr.y-1 >= 0 && image[curr.x][curr.y-1] == oldColor {
				q = append(q, Coordinate{
					x: curr.x,
					y: curr.y - 1,
				})
			}

			if curr.y+1 < len(image[0]) && image[curr.x][curr.y+1] == oldColor {
				q = append(q, Coordinate{
					x: curr.x,
					y: curr.y + 1,
				})
			}
		}
	}

	return image
}
