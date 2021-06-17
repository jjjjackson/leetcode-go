package leetcode

// https://leetcode.com/problems/number-of-islands/description/
// 用 DFS 或 BFS 看是不是同一個 Island
// 然後總共有幾個 Island
// 這裡用 DFS

type Coordinate struct {
	X, Y int
}

var directions = [4]Coordinate{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

const SEA = '0'
const ISLAND = '1'

func numIslands(grid [][]byte) int {
	numOfIslands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == ISLAND {
				exploreIsland(grid, Coordinate{i, j})
				numOfIslands++
			}
		}
	}
	return numOfIslands
}

func exploreIsland(grid [][]byte, from Coordinate) {
	stack := []Coordinate{from}

	for len(stack) > 0 {
		coord := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if coord.X < 0 || coord.Y < 0 || coord.X >= len(grid) || coord.Y >= len(grid[0]) {
			continue
		}

		if grid[coord.X][coord.Y] == SEA {
			continue
		}

		grid[coord.X][coord.Y] = SEA

		for _, d := range directions {
			stack = append(stack, Coordinate{d.X + coord.X, d.Y + coord.Y})
		}
	}
}
