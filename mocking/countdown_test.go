package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	
	t.Run("sleep 4 times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if spySleeper.calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.calls)
		}
	})

	t.Run("sleep after print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleeperPrinter.calls) {
			t.Errorf("got '%v' want '%v'", spySleeperPrinter.calls, want)
		}

	})
}

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep()  {
	s.calls++
}

//既实现了 sleeper 又实现了 writer
type CountdownOperationSpy struct {
	calls []string
}

const sleep = "sleep"
const write = "write"
func (s *CountdownOperationSpy) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error)  {
	s.calls = append(s.calls,  write)
	return 
}
