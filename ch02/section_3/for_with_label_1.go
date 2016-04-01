package main

import "fmt"

func main() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				// LOOP로 지정된 for문을 빠져나옴
				break LOOP
			}

		}
	}
}
