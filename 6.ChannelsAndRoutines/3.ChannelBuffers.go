//By Default channels are unbuffered
//i.e they will accept if corresponding is reacdy to send
//Buffered channels accept a limited number of values
//without a corresponding receiver for those values.

package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
