package main

import (
	"fmt"
	"sync"
)

var balance = 1000
var rw sync.RWMutex

func readBalance() {
	rw.RLock()

	fmt.Println("Balance:", balance)

	rw.RUnlock()
}

func main() {

	readBalance()
	readBalance()
	readBalance()
}