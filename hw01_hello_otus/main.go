package main

import (
	"fmt"
	utils "github.com/leighmcculloch/go-strrev"
)

func main() {
	var s string
	s = "Hello, OTUS!"
	fmt.Println(utils.Reverse(s))
}
