package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord  = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}
//生产环境是 1 秒， 测试环境统计调用次数
func main(){
	sleeper := &ConfigurableSleeper{ 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
