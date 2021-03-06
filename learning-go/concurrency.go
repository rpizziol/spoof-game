package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { // main go routine
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()

	//go count("sheep")
	//go count("fish")
	//fmt.Scanln()
	//time.Sleep(time.Millisecond * 2000)
}

/*func main() {
	go count("fish")
	count("sheep")
}*/

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
