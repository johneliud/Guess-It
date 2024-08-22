package main

import (
	"fmt"
	"os"

	"github.com/johneliud/Guess-It/readstandardinput"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	readstandardinput.ReadStandardInput()
}
