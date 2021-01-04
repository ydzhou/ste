package main

import (
	"fmt"
	"github.com/ydzhou/ste/internal"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("failed to start: too many arguments\nusage: ste filename")
	}

	editor := ste.Editor{}
	editor.Init()

	if len(args) > 0 {
		editor.Open(args[0])
	}

	editor.Start()
}
