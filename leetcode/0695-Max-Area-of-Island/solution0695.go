package leetcode

// https://leetcode.com/problems/max-area-of-island/submissions/
// 找最大區域
// DFS BFS 都可
// 這邊用 DFS

type Coordinate struct {
	X, Y int
}

var directions = [4]Coordinate{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 1 {
				result = max(result, exploreIsland(grid, Coordinate{x, y}))
			}
		}
	}
	return result
}

func exploreIsland(grid [][]int, from Coordinate) int {
	stack := []Coordinate{from}
	size := 0

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !inRange(grid, p) {
			continue
		}

		if grid[p.X][p.Y] == 0 {
			continue
		}

		grid[p.X][p.Y] = 0
		size++

		for _, d := range directions {
			n := Coordinate{d.X + p.X, d.Y + p.Y}
			stack = append(stack, n)
		}
	}

	return size
}

func inRange(grid [][]int, p Coordinate) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < len(grid) && p.Y < len(grid[len(grid)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
