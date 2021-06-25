package leetcode

//BFS 或 DP 都可以解
// DP => f[x,y] = min{ f[x-1,y], f[x, y-1] } + grid[x][y]
// Boundary => f[0,0] = grid[0][0]
// Boundary => x-1 || y-1 <0 => take another
// 可以用滾動數組優化

func minPathSum(grid [][]int) int {
	// Runtime 8ms
	// Memory 3.9MB

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	lenX := len(grid)
	lenY := len(grid[0])
	dp := make([][]int, 2)
	prevRow, curRow := 0, 1

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, lenY)
	}

	for x := 0; x < lenX; x++ {
		prevRow, curRow = curRow, prevRow

		for y := 0; y < lenY; y++ {

			if x == 0 && y == 0 {
				dp[curRow][0] = grid[0][0]
				continue
			}

			if x-1 < 0 {
				dp[curRow][y] = dp[curRow][y-1] + grid[x][y]
				continue
			}

			if y-1 < 0 {
				dp[curRow][y] = dp[prevRow][y] + grid[x][y]
				continue
			}

			dp[curRow][y] = min(dp[curRow][y-1], dp[prevRow][y]) + grid[x][y]

		}
	}

	return dp[curRow][lenY-1]
}

func minPathSum1(grid [][]int) int {
	// Runtime 4ms
	// Memory 4.4MB
	// 用空間換取時間

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	lenX := len(grid)
	lenY := len(grid[0])
	dp := make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		dp[i] = make([]int, lenY)
	}

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {

			if x == 0 && y == 0 {
				dp[0][0] = grid[0][0]
				continue
			}

			if x-1 < 0 {
				dp[x][y] = dp[x][y-1] + grid[x][y]
				continue
			}

			if y-1 < 0 {
				dp[x][y] = dp[x-1][y] + grid[x][y]
				continue
			}

			dp[x][y] = min(dp[x][y-1], dp[x-1][y]) + grid[x][y]

		}
	}

	return dp[lenX-1][lenY-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
