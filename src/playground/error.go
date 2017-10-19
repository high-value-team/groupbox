package main

import (
	"errors"
	"fmt"
)

var SadError = errors.New("Something really sad happened")

func main() {
	var err error
	a(&err)
	b(&err)
	fmt.Printf("Error:%+v", err)
}

func a(err *error) {
	if *err != nil {
		return
	}
	*err = SadError
	//*err = errors.New("Mein Error A!")
	//*err = fmt.Errorf("Mein Error A!")
}

func b(err *error) {
	if *err != nil {
		return
	}
	*err = fmt.Errorf("Mein Error B!")
}
