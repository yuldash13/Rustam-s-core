package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var mapa = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}

	//for i := 0; i < 100; i++ {
	//	go func() {
	//		consumer(mapa)
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	j := i
	//	go func() {
	//		producer(mapa, j)
	//	}()
	//}
	//time.Sleep(10 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	rw := sync.RWMutex{}

	go func() {
		counter(cancel)
	}()
	producer(mapa, &rw, ctx)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			consumer(mapa, &rw)
		}
	}
}

func counter(cancel context.CancelFunc) {
	mu := sync.Mutex{}

	var count int

	for {
		if count == 100 {
			cancel()
			return
		}
		mu.Lock()
		count++
		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func producer(mapa map[int]int, rw *sync.RWMutex, ctx context.Context) {
	var n = 1
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				if i == len(mapa) {
					i = 0
					n++
				}
				rw.Lock()
				mapa[i] = n
				rw.Unlock()
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func consumer(mapa map[int]int, rw *sync.RWMutex) {
	for i := 0; i < 100; i++ {
		go func() {
			rw.RLock()
			fmt.Println(mapa)
			rw.RUnlock()
		}()
	}
}
