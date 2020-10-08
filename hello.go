package main

import (
	"fmt"
	"time"
)

func main() {

	// for i := 0; i < 20; i++ {
	// 	fmt.Println("Hello world")
	// }

	go fmt.Println("Hello from goroutine 0")
	go fmt.Println("Hello from goroutine 1")
	go fmt.Println("Hello from goroutine 2")
	go fmt.Println("Hello from goroutine 3")
	go fmt.Println("Hello from goroutine 4")

	time.Sleep(1 * time.Second)
}
