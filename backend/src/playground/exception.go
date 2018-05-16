package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("before a\n")

	//go a() // panic immediatly -> terminates program execution
	go b() // panic after main finishes -> NO termination

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("after a\n")
}

func a() {
	panic("in a")
}

func b() {
	time.Sleep(200 * time.Millisecond)
	panic("in b")
}
