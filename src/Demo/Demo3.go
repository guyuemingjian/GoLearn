package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go PrintHello(i,ch)
	}
	for {
		msg := <- ch
		fmt.Println(msg)
	}

	//time.Sleep(100 * time.Millisecond)
}

func PrintHello(wkCount int,wkCh chan string) {
	for {
		wkCh <- fmt.Sprintf("hello world %d \n", wkCount)
	}

}
