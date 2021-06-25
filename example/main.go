package main

import (
	"fmt"
	"github.com/zhaohuabing/debounce"
	"time"
)
func main() {
	startTime:=time.Now()
	callback := func() {
		duration:=time.Since(startTime)
		startTime=time.Now()
		fmt.Printf("duration: %v\n",duration)
	}
	stop := make(chan struct{})
	d := debounce.New(200*time.Millisecond, 1*time.Second, callback, stop)
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		d.Bounce()
		if i== 10{
			d.Cancel()
		}
	}
	stop <- struct{}{}
}
