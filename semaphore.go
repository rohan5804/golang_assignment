package main

import (
	"fmt"
	"time"
)

func main() {

	sem := make(chan struct{}, 3)

	for i := 1; i <= 10; i++ {

		sem <- struct{}{}

		go func(id int) {

			fmt.Println("Working", id)

			time.Sleep(time.Second)

			fmt.Println("Done", id)

			<-sem

		}(i)
	}

	time.Sleep(5 * time.Second)
}