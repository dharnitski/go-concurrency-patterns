package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				// code here may be not executed
			}
		}
	}()
	return c // Return the channel to the caller
}

func cleanup(){
	fmt.Println("Cleanup")
}
