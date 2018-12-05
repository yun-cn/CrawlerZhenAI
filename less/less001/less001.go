package main

import "fmt"

func main() {

}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pupm2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
