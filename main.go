package main

import "fmt"

func main() {
	channel := make(chan any)
	go GorotineDoubleIt(channel, 23)
	fmt.Println(<-channel)
}
