package synctimer

import (
	"sync"
	"time"
)

type Timer struct {
	syncChannel chan bool
	wg          sync.WaitGroup
}

func NewTimer() *Timer {
	return &Timer{
		syncChannel: make(chan bool),
		wg:          sync.WaitGroup{},
	}
}

func (timer *Timer) NewSubTimer(duration time.Duration) chan bool {
	c := make(chan bool, 1)
	timer.wg.Add(1)
	go func() {
		<-timer.syncChannel
		time.Sleep(duration)
		close(c)
		timer.wg.Done()
	}()
	return c
}

func (timer *Timer) Start() {
	close(timer.syncChannel)
	timer.wg.Wait()
}
