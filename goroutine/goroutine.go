package goroutine

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		milliSecond := time.Millisecond
		//fmt.Println(milliSecond)
		time.Sleep(100 * milliSecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
