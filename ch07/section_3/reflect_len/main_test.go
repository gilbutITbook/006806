package main

import (
	"strconv"
	"testing"
)

func TestLenForMap(t *testing.T) {
	v := map[string]int{"A": 1, "B": 2}
	actual, expected := Len(v), 2
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
func TestLenForString(t *testing.T) {
	v := "one"
	actual, expected := Len(v), 3
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
func TestLenForSlice(t *testing.T) {
	v := []int{5, 0, 4, 1}
	actual, expected := Len(v), 4
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

/*
func TestLenForChan(t *testing.T) {
	v := make(chan int)
	actual, expected := Len(v), 1
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
*/

func BenchmarkLenForString(b *testing.B) {
	b.StopTimer()
	v := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		v[i] = strconv.Itoa(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Len(v[i%1000000])
	}
}
