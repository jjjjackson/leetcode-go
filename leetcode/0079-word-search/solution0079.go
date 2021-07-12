package leetcode

// https://leetcode.com/problems/word-search/
// DFS

type Coordinate struct {
	X, Y int
}
type Visited map[Coordinate]bool

func (c Coordinate) neighbors() []Coordinate {
	directions := [4]Coordinate{
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 0},
	}

	coords := make([]Coordinate, len(directions))

	for i, d := range directions {
		coords[i] = Coordinate{
			c.X + d.X,
			c.Y + d.Y,
		}
	}

	return coords
}

func (c Coordinate) isOutOfBoarder(lenX, lenY int) bool {
	return c.X < 0 || c.X >= lenX || c.Y < 0 || c.Y >= lenY
}

func isFoundWord(board [][]byte, word string, visited Visited, coord Coordinate, wordPointer int) bool {

	if wordPointer == len(word)-1 {
		return board[coord.X][coord.Y] == word[wordPointer]
	}

	if board[coord.X][coord.Y] != word[wordPointer] {
		return false
	}

	visited[coord] = true
	for _, n := range coord.neighbors() {
		if v, ok := visited[n]; ok && v {
			continue
		}
		if n.isOutOfBoarder(len(board), len(board[0])) {
			continue
		}
		if isFoundWord(board, word, visited, n, wordPointer+1) {
			return true
		}
	}

	visited[coord] = false

	return false
}

func exist(board [][]byte, word string) bool {

	if board == nil || board[0] == nil {
		return false
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	boardXLen := len(board)
	boardYLen := len(board[0])

	for i := 0; i < boardXLen; i++ {
		for j := 0; j < boardYLen; j++ {

			if board[i][j] == word[0] {
				visited := make(Visited)
				if isFoundWord(board, word, visited, Coordinate{i, j}, 0) {
					return true
				}
			}
		}
	}

	return false
}
