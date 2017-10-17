package main

import (
	"github.com/mjibson/esc/embed"
)

func main() {
	embed.Run(&embed.Config{
		Private:    true,
		Package:    "frontend",
		OutputFile: "frontend/frontend.go",
		Prefix:     "frontend/build",
		Files:      []string{"frontend/build"},
	})
}
