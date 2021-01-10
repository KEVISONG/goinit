package main

import (
	"fmt"
	"os"

	"github.com/KEVISONG/goinit/pkg/entry"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please give your project a name :)")
		return
	}

	entry.Init(os.Args[1])
}
