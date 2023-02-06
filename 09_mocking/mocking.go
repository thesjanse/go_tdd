package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1)
}

func Countdown(writer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		s.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
