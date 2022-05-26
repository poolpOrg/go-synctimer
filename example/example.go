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

	wg.Add(1)
	c1 := t.NewSubTimer(5 * time.Second)
	go func() {
		<-c1
		fmt.Println("wake up #1!")
		wg.Done()
	}()

	wg.Add(1)
	c2 := t.NewSubTimer(5 * time.Second)
	go func() {
		<-c2
		fmt.Println("wake up #2!")
		wg.Done()
	}()

	wg.Add(1)
	c3 := t.NewSubTimer(6 * time.Second)
	go func() {
		<-c3
		fmt.Println("wake up #3!")
		wg.Done()
	}()

	t.Start()
	wg.Wait()
}
