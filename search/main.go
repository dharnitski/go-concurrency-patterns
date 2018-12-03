package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
	Web2   = fakeSearch("web")
	Image2 = fakeSearch("image")
	Video2 = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// sync version - run all searches sequentially
// takes up to 300ms
// func Google(query string) (results []Result) {
//     results = append(results, Web(query))
//     results = append(results, Image(query))
//     results = append(results, Video(query))
//     return
// }


// async version
// takes up to 100ms
// func Google(query string) (results []Result) {
// 	c := make(chan Result)
// 	go func() { c <- Web(query) }()
// 	go func() { c <- Image(query) }()
// 	go func() { c <- Video(query) }()

// 	for i := 0; i < 3; i++ {
// 		result := <-c
// 		results = append(results, result)
// 	}
// 	return
// }


// async version - do not wait for long responses
// returns only results that take less than 80ms
// func Google(query string) (results []Result) {
// 	c := make(chan Result)
// 	go func() { c <- Web(query) }()
// 	go func() { c <- Image(query) }()
// 	go func() { c <- Video(query) }()

// 	timeout := time.After(80 * time.Millisecond)
// 	for i := 0; i < 3; i++ {
// 		select {
// 		case result := <-c:
// 			results = append(results, result)
// 		case <-timeout:
// 			fmt.Println("timed out")
// 			return
// 		}
// 	}
// 	return
// }

// redundant 
func Google(query string) (results []Result) {
    c := make(chan Result)
    go func() { c <- First(query, Web, Web2) } ()
    go func() { c <- First(query, Image, Image2) } ()
    go func() { c <- First(query, Video, Video2) } ()
    timeout := time.After(80 * time.Millisecond)
    for i := 0; i < 3; i++ {
        select {
        case result := <-c:
            results = append(results, result)
        case <-timeout:
            fmt.Println("timed out")
            return
        }
    }
    return
}

func First(query string, replicas ...Search) Result {
    c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	// run function multiple times in separate goroutines
    for i := range replicas {
        go searchReplica(i)
	}
	// wait for first result
    return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

