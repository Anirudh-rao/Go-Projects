//Channels are the pipes that connect concurrent goroutines
package main

import (
	"fmt"
)

func main() {
	//Creates a Chan named messages
	messages := make(chan string)

	//Routine to recive channel indicated by <- sign
	go func() { messages <- "ping" }()

	// <- is  used to indicate reciving the channel
	msg := <-messages
	fmt.Println(msg)
}
