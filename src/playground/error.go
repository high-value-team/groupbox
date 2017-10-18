package main

import "fmt"

func main() {
	var err error
	a(&err)
	fmt.Printf("Error:%+v", err)
}

func a(err *error) {
	if *err != nil {
		return
	}
	*err = fmt.Errorf("Mein Error!")
}
