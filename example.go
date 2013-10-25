package main

import "fmt"
import "time"

import "github.com/nickdavies/go-ratelimit/ratelimit"

func main () {
    limiter := ratelimit.New(5, time.Second, false)

    start := time.Now()
    events := 0
    for {
        fmt.Printf("Time Past %s Events sent %d\n", time.Now().Sub(start), events)

        if err := limiter.Event(); err != nil {
            fmt.Println("Error: ", err)
            break
        }
        fmt.Println("Event")
        events++

        time.Sleep(100 * time.Millisecond)
    }
}

