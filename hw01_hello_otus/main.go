package main

import (
	"fmt"
	utils "github.com/leighmcculloch/go-strrev"
)

func main() {
	s := "Hello, OTUS!"
	fmt.Println(utils.Reverse(s))
}
