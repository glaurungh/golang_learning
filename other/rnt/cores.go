package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go printNums(10)
	runtime.Gosched()
	fmt.Println("done")
}

func printNums(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
