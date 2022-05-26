# go-synctimer

synctimer is a package that provides synchronized timers.

Golang's  `time.Timer` is given a `time.Duration` relative to `time.Now()` starts counting right away.

This makes it unsuitable for cases where the timer needs to be prepared in advance but only start its countdown when a condition is met,
or for cases where a collection of timers need to be synchronized so that they begin their countdown simultaneously (ie: if we set timers that are supposed to start relative to each other and only begin the coundown when the last one is configured).

The synctimer package allows creating a `synctimer.Timer` object,
in charge of synchronizing a set of subcounters,
launching them all simultaneously:

```go
func main() {
    // this is just so the example doesn't exit right after t.Start()
    // and before all fmt.Println() calls are visible
	wg := sync.WaitGroup{}

    // create a new synchronized timer
	t := synctimer.NewTimer()

    // create ten subtimers, all waking up at 1 up-to 10 seconds
    // relative to the synchronized t.Start()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int, c chan bool) {
			<-c
			fmt.Println("woke up subtimer", id)
			wg.Done()
		}(i, t.NewSubTimer(time.Duration(i)*time.Second))
	}

    // launch all timers !
	t.Start()

    // wait for all local goroutines to be done so we know that the
    // example worked :-)
	wg.Wait()
}
```