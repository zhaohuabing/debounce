# debouncer
Debouncer ensures a method call will not be executed too frequently. Debouncer only call the callback function after a
 given time without new bounce event, or the time after the previous execution reaches the max duration.

# Usage

```go
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
```



