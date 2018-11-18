package hourglass

import "fmt"

// MaxHourGlass finds the hourglass in the grid that gives
// the largest sum.
func MaxHourGlass(arr [][]int) string {
	height := len(arr)
	width := len(arr[0])
	maxSum := 0
	var maxI, maxJ int
	for i := 0; i < height-2; i++ {
		for j := 0; j < width-2; j++ {
			sum := HourGlassSum(arr, i, j)
			if sum > maxSum {
				maxSum = sum
				maxI, maxJ = i, j
			}
		}
	}
	return PrintHourGlass(arr, maxI, maxJ)
}

// HourGlassSum returns the sum of the particular hourglass
// with top-left element (i, j)
func HourGlassSum(arr [][]int, i, j int) int {
	return arr[i][j] + arr[i][j+1] + arr[i][j+2] +
		arr[i+1][j+1] +
		arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
}

// PrintHourGlass prints out the hourglass at the given coordinates
func PrintHourGlass(arr [][]int, i, j int) string {
	return fmt.Sprintf("%d %d %d\n  %d\n%d %d %d",
		arr[i][j], arr[i][j+1], arr[i][j+2],
		arr[i+1][j+1],
		arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])
}
