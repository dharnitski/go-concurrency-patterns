package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generator for two persons
// blocks each other
// func main() {
// 	joe := boring("Joe")
// 	ann := boring("Ann")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-joe)
// 		fmt.Println(<-ann)
// 	}
// 	fmt.Println("You're boring; I am leaving.")
// }

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I am leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}
