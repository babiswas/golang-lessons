package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("hello World")
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("hello bello world")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second)
	fmt.Scanln()
}
