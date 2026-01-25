package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if m == 0 || n == 0 {
		return 0
	}

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for col := range m {
		dp[col] = make([]int, n)
	}

	dp[0][0] = 1
	for row := 1; row < m; row++ {
		if obstacleGrid[row][0] == 1 {
			dp[row][0] = 0 // obstaculo
		} else {
			dp[row][0] = dp[row-1][0]
		}
	}

	// for col := range n {
	for col := 1; col < n; col++ {
		if obstacleGrid[0][col] == 1 {
			dp[0][col] = 0 // obstaculo
		} else {
			dp[0][col] = dp[0][col-1]
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0
			} else {
				dp[row][col] = dp[row-1][col] + dp[row][col-1]
			}
		}
	}

	return dp[m-1][n-1]
}
