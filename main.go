package main

import (
	"fmt"
	"os"
)

func main() {
	//	for i := range os.Args {
	//		fmt.Println("arg: ", i, os.Args[i])
	//	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	data := make([]byte, 32*1024)
	count, errF := file.Read(data)
	if errF != nil {
		fmt.Println("Error: ", errF)
		os.Exit(1)
	}
	fmt.Println(string(data[:count]))
	//	io.Copy(os.Stdout, )

}
