package interactions

import "fmt"

func async(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Exception occured: %+v", r)
		}
	}()
	go fn()
}
