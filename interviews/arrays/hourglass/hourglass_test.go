package hourglass

import "testing"

var grid = [][]int{
	{1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0},
	{0, 0, 2, 4, 4, 0},
	{0, 0, 0, 2, 0, 0},
	{0, 0, 1, 2, 4, 0},
}

func TestMaxHourGlass(t *testing.T) {
	want := `2 4 4
  2
1 2 4`
	got := MaxHourGlass(grid)
	if got != want {
		t.Errorf("\nMax hourglass:\n%s\nWant:\n%s\n", got, want)
	}
}

func TestHourGlassSum(t *testing.T) {
	want := 19
	got := HourGlassSum(grid, 3, 2)
	if got != want {
		t.Errorf("Sum of hourglass = %d; expected %d", got, want)
	}
}
