package leetcode

func Sol_medium_200(grid [][]byte) int {
	var dfs func(grid [][]byte, i, j int)

	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}

	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/
////
// 1. Iterate through each cell in the 2D grid
// 2. If we encounter a land cell, we increment the island count
// 3. Once we find a land cell, we perform DFS to mark all connected land cells as visited ('0')
// 4. Repeat step 2 and 3 until we visit all cells in the grid
// 5. Return the island count
