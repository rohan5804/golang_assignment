package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	fmt.Println(n)
}

func main() {

	for i := 1; i <= 5; i++ {
		go printNumber(i)
	}

	time.Sleep(2 * time.Second)
}
