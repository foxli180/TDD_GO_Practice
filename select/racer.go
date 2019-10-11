package _select

import (
	"fmt"
	"net/http"
	"time"
)

//func Racer(a, b string) (winner string) {
//
//	aDuration := measureResponseTime(a)
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//	return b
//}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, err
	case <-ping(b):
		return b, err
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

var tenSecondsTimeout = 10 * time.Second


func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
