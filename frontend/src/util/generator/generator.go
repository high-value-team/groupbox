package main

import (
	"fmt"

	"github.com/mjibson/esc/embed"
)

func main() {
	if err := embed.Run(&embed.Config{
		Private:    false,
		Package:    "frontend",
		OutputFile: "frontend/frontend.go",
		Prefix:     "frontend/build",
		Files:      []string{"frontend/build"},
	}); err != nil {
		fmt.Println(err)
	}
}
