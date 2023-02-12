package selecto

import (
	"errors"
	"net/http"
	"time"
)

var ErrorTimeout = errors.New("Response took more than 10 secords. Timeout")

const defaultDelay = 10 * time.Second

func Racer(a, b string) (url string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(defaultDelay):
		return "", ErrorTimeout
	}
}

func ConfigurableRacer(a, b string, timeout time.Duration) (url string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
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
