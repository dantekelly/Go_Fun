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
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, s Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
