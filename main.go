package main

import (
	"fmt"
	"os"
)

func main() {
	for i := range os.Args {
		fmt.Println("arg: ", i, os.Args[i])
	}
}
