package leetcode

// https://leetcode.com/problems/surrounded-regions/description/
// DFS
// 將邊緣的 O 標記成另一種符號
// 然後再將 O 都表示為 X

type Coordinate struct {
	X, Y int
}

var directions = [4]Coordinate{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

const BOARD = 'X'
const SPACE = 'O'
const TERRITORY = 'T'

func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == SPACE && isBoader(board, Coordinate{i, j}) {
				dfs(board, Coordinate{i, j})
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == TERRITORY {
				board[i][j] = SPACE
			} else if board[i][j] == SPACE {
				board[i][j] = BOARD
			}
		}
	}
}

func dfs(board [][]byte, coord Coordinate) {
	if isOutOfRange(board, coord) {
		return
	}

	if board[coord.X][coord.Y] == SPACE {
		board[coord.X][coord.Y] = TERRITORY
		for _, d := range directions {
			dfs(board, Coordinate{coord.X + d.X, coord.Y + d.Y})
		}
	}
}

func isOutOfRange(board [][]byte, coord Coordinate) bool {
	if coord.X < 0 || coord.Y < 0 || coord.X >= len(board) || coord.Y >= len(board[len(board)-1]) {
		return true
	}
	return false
}

func isBoader(board [][]byte, coord Coordinate) bool {
	if coord.X == 0 || coord.Y == 0 || coord.X == len(board)-1 || coord.Y == len(board[len(board)-1])-1 {
		return true
	}
	return false
}
