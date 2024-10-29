package main

import (
	"fmt"
	"time"
)

func Sleep(ms int) {
	<-time.After(time.Duration(ms) * time.Millisecond)
}
func main() {
	// time.Sleep(10 * time.Millisecond)
	fmt.Println("Start")
	Sleep(2000)
	fmt.Println("Stop")
}