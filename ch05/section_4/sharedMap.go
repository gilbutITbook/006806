package main

import "fmt"

type sharedMap struct {
	m map[string]interface{}
	c chan command
}

type command struct {
	action int
	key    string
	value  interface{}
	result chan<- interface{}
}

const (
	set = iota
	get
	remove
	count
)

func (sm sharedMap) Set(k string, v interface{}) {
	sm.c <- command{action: set, key: k, value: v}
}

func (sm sharedMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (sm sharedMap) Remove(k string) {
	sm.c <- command{action: remove, key: k}
}

func (sm sharedMap) Count() int {
	callback := make(chan interface{})
	sm.c <- command{action: count, result: callback}
	return (<-callback).(int)
}

func (sm sharedMap) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

func NewMap() sharedMap {
	sm := sharedMap{
		m: make(map[string]interface{}),
		c: make(chan command),
	}
	go sm.run()
	return sm
}

func main() {
	m := NewMap()

	// Set item
	m.Set("foo", "bar")

	// Get item
	t, ok := m.Get("foo")

	// Check if item exists
	if ok {
		bar := t.(string)
		fmt.Println("bar: ", bar)
	}

	// Remove item
	m.Remove("foo")

	// Count
	fmt.Println("Count: ", m.Count())
}
