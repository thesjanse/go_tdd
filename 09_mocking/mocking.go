package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
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
