package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for n, arg := range os.Args[1:] {
		fmt.Printf("%d\t%v\n", n, arg)
	}
	fmt.Println(s)
}
