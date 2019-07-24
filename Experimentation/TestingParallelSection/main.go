package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

func main() {
	count = 0
	var wg sync.WaitGroup
	fmt.Println("working")
	wg.Add(1)
	go func() {
		fmt.Println("func one")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("func two")
		wg.Done()
	}()
	//fmt.Scanln()
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go updateCount(&wg, &mu)
	}
	wg.Wait()
	fmt.Println(count)
}

var happenOnce bool

func updateCount(wg *sync.WaitGroup, mu *sync.Mutex) {
	fmt.Println("starting func")
	for i := 0; i < 100; i++ {
		mu.Lock()
		if !happenOnce && i == 100 {
			happenOnce = true
			wg.Done()
			fmt.Println("ending the one time ender")
			return
		}
		count++
		mu.Unlock()
		runtime.Gosched()
	}
	fmt.Println("ending func")
	wg.Done()
}
