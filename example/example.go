package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/poolpOrg/go-synctimer"
)

func main() {
	wg := sync.WaitGroup{}

	t := synctimer.NewTimer()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int, c chan bool) {
			<-c
			fmt.Println("woke up subtimer", id)
			wg.Done()
		}(i, t.NewSubTimer(time.Duration(i)*time.Second))
	}
	t.Start()
	wg.Wait()
}
