package selecto

import (
	"errors"
	"net/http"
	"time"
)

var ErrorTimeout = errors.New("Response took more than 10 secords. Timeout")

func Racer(a, b string) (url string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", ErrorTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
