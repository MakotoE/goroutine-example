package main

import (
	"fmt"
	"time"
)

func main() {
	go world()

	fmt.Println("Hello")
	time.Sleep(time.Second)
}

func world() {
	fmt.Println("world")
}
