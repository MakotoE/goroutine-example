package main

import (
	"fmt"
)

func main() {
	stringChannel := make(chan string)

	go func() {
		fmt.Println("world")
		stringChannel <- "!"
	}()

	fmt.Println("Hello")
	result := <-stringChannel

	fmt.Println(result)
}
