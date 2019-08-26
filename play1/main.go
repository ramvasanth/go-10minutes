package main

import (
	"runtime"
	"sync"
)

var numbers = map[string]int{"1": 1, "2": 2, "3": 3}
var urls = []string{"1", "2", "3", "4"}

func mockGet(url string) int {
	if v, ok := numbers[url]; ok {
		return v
	}

	return 0
}

func useMutex() int {
	bigNumber := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(urls))
	get := func(url string) {
		n := mockGet(url)
		mu.Lock()
		if n > bigNumber {
			bigNumber = n
		}
		mu.Unlock()
		wg.Done()
	}

	for i := range urls {
		go get(urls[i])
	}
	wg.Wait()
	return bigNumber
}
func main() {
	runtime.CPUProfile()
}

func useChannnel() int {
	ch := make(chan int, len(urls))

	for i := range urls {
		go visitURL(urls[i], ch)
	}

	bigNumber := 0
	for i := 0; i < len(urls); i++ {
		n := <-ch
		if n > bigNumber {
			bigNumber = n
		}
	}
	return bigNumber
}
func visitURL(url string, ch chan int) {
	n := mockGet(url)
	ch <- n
}
