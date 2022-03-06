package main

import (
	"fmt"
	"time"
)

func main() {

	//Making two channels c1
	c1 := make(chan string)
	c2 := make(chan string)

	//Go routine for c1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	//Go routine for c2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		//Selection of channels
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
