package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// withoutWait()
	// withWait()

	// writeWithoutConcurrent()
	// writeWithoutMutex()
	// writeWithMutex()
	readWithMutex()
	readWithRWMutex()

}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	fmt.Println("finish")
}

func withWait() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}

func writeWithoutConcurrent() {
	start := time.Now()

	counter := 0

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond * 10000)
		counter++
	}
	fmt.Println(counter)

	fmt.Printf("Time took: %.6f s\n", time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()

	counter := 0

	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond * 10000)
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)

	fmt.Printf("Time took: %.6f s\n", time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()

	counter := 0

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(m *sync.Mutex) {
			defer wg.Done()
			time.Sleep(time.Nanosecond * 10000)
			m.Lock()
			defer m.Unlock()
			counter++
		}(&m)
	}

	wg.Wait()
	fmt.Println(counter)

	fmt.Printf("Time took: %.6f s\n", time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			time.Sleep(time.Nanosecond)
			_ = counter
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			time.Sleep(time.Nanosecond)
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mu.RLock()
			defer mu.RUnlock()

			time.Sleep(time.Nanosecond)
			_ = counter
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			time.Sleep(time.Nanosecond)
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
