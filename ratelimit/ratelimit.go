package ratelimit

import (
    "errors"
    "time"
)

var RateExceededErr = errors.New("Error rate was exceeded")

type RateLimiter interface {
    Event() (err error)
}

func NewRateLimiter(limit int, panic_on_limit bool) RateLimiter {
    er := &rateLimit{
        events: make(chan interface{}, limit),
        current: 0,
        limit: limit,
        panic_on_limit: panic_on_limit,
    }
    go er.drain()

    return er
}

// Private Implementation

type rateLimit struct {
    events chan interface{}
    current int
    limit int
    panic_on_limit bool
}

func (e *rateLimit) Event() (err error) {
    select {
    case e.events <-nil:
        return nil
    default:
        if e.panic_on_limit {
            panic(RateExceededErr)
        }
        return RateExceededErr
    }
}

func (e *rateLimit) drain() {
    for _ = range time.Tick(time.Second / time.Duration(e.limit)) {
        select {
        case <-e.events:
            continue
        default:
            continue
        }
    }
}
