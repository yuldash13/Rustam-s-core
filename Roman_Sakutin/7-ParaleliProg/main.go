package main

import (
	"errors"
	"fmt"
	"sync"
)

type Config struct {
	FactorizationWorkers int
	WriteWorkers         int
}

func main() {
	arr := []int{100, -17, 25, 38}
	n := 3
	cfg := Config{
		FactorizationWorkers: len(arr),
		WriteWorkers:         n,
	}
	done := make(chan struct{})

	err := Do(done, arr, &cfg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("done")
	}
}

func Do(done chan struct{}, arr []int, cfg *Config) error {
	var (
		ErrFactorizationCancelled = errors.New("cancelled loser")
		ErrWriterInteraction      = errors.New("writer interaction loser")
		err                       = make(chan error)
	)

	jobs := make(chan int)
	go func() {
		for _, a := range arr {
			select {
			case <-done:
				close(jobs)
				return
			case jobs <- a:
			}
		}
		defer close(jobs)
	}()

	result := make(chan []int)

	Fwg := sync.WaitGroup{}

	for i := 0; i < cfg.FactorizationWorkers; i++ {
		Fwg.Add(1)
		go func() {
			defer Fwg.Done()
			for j := range jobs {
				select {
				case <-done:
					err <- ErrFactorizationCancelled
					return
				case result <- factor(j):
				}
			}
		}()
	}
	go func() {
		Fwg.Wait()
		close(result)
	}()

	Wwg := sync.WaitGroup{}

	for i := 0; i < cfg.WriteWorkers; i++ {
		Wwg.Add(1)
		go func() {
			defer Wwg.Done()
			for r := range result {
				select {
				case <-done:
					err <- ErrWriterInteraction
					return
				default:
					str := consumer(r)
					fmt.Println(str)
				}
			}
		}()
	}
	go func() {
		Wwg.Wait()
		close(err)
	}()

	return <-err
}

func factor(n int) []int {
	arr := make([]int, 0)
	arr = append(arr, n)
	if n == 1 {
		arr = append(arr, 1)
	} else if n == 0 {
		arr = append(arr, 0)
	} else if n < 0 {
		arr = append(arr, -1)
		n *= -1
	}
	for i := 2; n >= 2; i++ {
		if n%i == 0 {
			n /= i
			arr = append(arr, i)
			i--
		}
	}
	return arr
}

func consumer(arr []int) string {
	var str string
	str += fmt.Sprintf("%d = ", arr[0])
	for i := 1; i < len(arr); i++ {
		str += fmt.Sprintf("%d", arr[i])
		if i != len(arr)-1 {
			str += fmt.Sprintf("*")
		}
	}
	str += fmt.Sprintf("\n")
	return str
}
