package leetcode

// https://leetcode.com/problems/shortest-path-in-binary-matrix/

type Coordinate struct {
	X, Y int
}

const MAX_INT = int(1<<32 - 1)

func shortestPathBinaryMatrix(grid [][]int) int {
	// BFS

	if grid == nil {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	if grid[0][0] == 1 {
		return -1
	}

	if grid[m-1][n-1] == 1 {
		return -1
	}

	directions := [8]Coordinate{
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	stack := []Coordinate{}
	stack = append(stack, Coordinate{0, 0})
	layer := 0

	for len(stack) != 0 {
		layer += 1
		layerSize := len(stack)

		for layerSize > 0 {
			layerSize--

			s := stack[0]
			stack = stack[1:]

			if grid[s.X][s.Y] == 1 {
				continue
			}

			if s.X == m-1 && s.Y == n-1 {
				return layer
			}

			grid[s.X][s.Y] = 1

			for _, d := range directions {
				x := s.X + d.X
				y := s.Y + d.Y

				if x < 0 || y < 0 || x >= m || y >= n {
					continue
				}
				stack = append(stack, Coordinate{x, y})
			}

		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
