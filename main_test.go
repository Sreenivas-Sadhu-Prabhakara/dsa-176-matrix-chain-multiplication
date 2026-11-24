package main

import "testing"

func TestMinMultiplications(t *testing.T) {
	cases := []struct {
		dims []int
		want int
	}{
		{[]int{10, 20, 30, 40}, 18000},
		{[]int{40, 20, 30, 10, 30}, 26000}, // classic CLRS-style case
		{[]int{10, 20, 30}, 6000},          // two matrices: one way only
		{[]int{5, 10}, 0},                  // single matrix: no multiplication
		{[]int{7}, 0},                      // degenerate
		// Both orders by hand: ((AB)C) = 2*3*4 + 2*4*3 = 48; (A(BC)) = 3*4*3 + 2*3*3 = 54.
		{[]int{2, 3, 4, 3}, 48},
	}
	for _, c := range cases {
		if got := MinMultiplications(c.dims); got != c.want {
			t.Errorf("MinMultiplications(%v) = %d, want %d", c.dims, got, c.want)
		}
	}
}
