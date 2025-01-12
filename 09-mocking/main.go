package _9_mocking

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (sleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (sleeper *ConfigurableSleeper) Sleep() {
	sleeper.sleep(sleeper.duration)
}

const finalWord = "Go!"

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(writer io.Writer, sleeper Sleeper, countdownStart int) {
	for i := range countDownFrom(countdownStart) {
		fmt.Fprintln(writer, strconv.Itoa(i))
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper, 3)
}
