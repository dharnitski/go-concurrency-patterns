# Code for Google I/O 2012 - Go Concurrency Patterns

https://www.youtube.com/watch?v=f6kdp27TYZs
https://talks.golang.org/2012/concurrency.slide#37

* Basic: simple concurrency sample with one person
* Blocking: Two person blocking each other
* Generator: function that returns a channel
* Multiplexing (fan-in): function that takes multiple channels and pipes to one channel, so that the returned channel receives both outputs
* Restoring: restoring sequencing after multiplexing
* Select: fanIn() using select
* Timeout message: stops execution if person is not responding 
* Timeout conversation: stops execution after certain period of time
* Quit: explicitly tell person to stop
* Cleanup: caller asks person to cleanup himself before quit
* Search: fake search engine

## how to run

    $ cd {$foldername}
    $ go run main.go