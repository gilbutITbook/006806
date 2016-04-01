package main

import "fmt"

func main() {
	groupMap := make(map[string]string)

	group1 := []int32{1, 4, 6}
	group2 := []int32{2, 4, 5}
	group3 := []int32{4, 6, 7}

	groupMap[string(group1)] = "first"
	groupMap[string(group2)] = "second"
	groupMap[string(group3)] = "third"

	fmt.Println(groupMap[string(group3)]) // third
	fmt.Println(groupMap[string(group2)]) // second
	fmt.Println(groupMap[string(group1)]) // first

}
