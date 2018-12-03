package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("exit")
}

func boring(msg string, quit <-chan bool) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// nothing
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()
	return c // Return the channel to the caller
}
