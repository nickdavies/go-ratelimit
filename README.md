go-ratelimit
============

A small library used for rate limiting events (such as errors or requests).

This uses a simplified version of the leaky bucket algorithm with the average rate limit
and maximum burst size set to the same value.

For example if you use the rate 5 per second you an error will be raised if you send on average more than 5 events per second
or if you try and send a burst of events of greater than 5. If you run the example you will see that you are not imediatly 
punished for exceeding the rate limit but only once you sustain the excessive rate.

Install with:

    go get github.com/nickdavies/go-ratelimit/ratelimit

Go Docs can be found at: http://godoc.org/github.com/nickdavies/go-ratelimit/ratelimit

Example:

```go
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
```

A running version of the example is in examle.go
