package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regexStr string

	path := "/api/boxes/1"
	format := " regexStr:%s\n path:%s\n match:%t\n\n\n"

	regexStr = "^/api/boxes$"
	fmt.Printf(format, regexStr, path, match(regexStr, path))

	regexStr = "^/api/boxes/([a-zA-Z0-9]+)$"
	fmt.Printf(format, regexStr, path, match(regexStr, path))
}

