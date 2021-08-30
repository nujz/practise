package main

func search(board [][]byte, word string, x int, y int, u int, points [][]bool) bool {
	if x < 0 || x >= len(board) {
		return false
	}
	if y < 0 || y >= len(board[x]) {
		return false
	}
	if board[x][y] != word[u] {
		return false
	}
	if points[x][y] {
		return false
	}
	if u == len(word)-1 {
		return true
	}

	points[x][y] = true
	if search(board, word, x-1, y, u+1, points) {
		return true
	}
	if search(board, word, x+1, y, u+1, points) {
		return true
	}
	if search(board, word, x, y-1, u+1, points) {
		return true
	}
	if search(board, word, x, y+1, u+1, points) {
		return true
	}
	points[x][y] = false

	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	var points = make([][]bool, len(board))
	for m := 0; m < len(board); m++ {
		points[m] = make([]bool, len(board[m]))
	}

	for m := 0; m < len(board); m++ {
		for n := 0; n < len(board[m]); n++ {
			if board[m][n] == word[0] {
				if search(board, word, m, n, 0, points) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESEEEFS"
	println(exist(board, word))
}
