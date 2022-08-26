package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Depositing", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdrawal(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Withdrawing", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	balance = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withdrawal(700, &wg)
	wg.Wait()

	fmt.Println("New balance", balance)

}
