package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			fmt.Println("Inside infinite loop")
			select {
			case <-done:
				fmt.Println("Ticker is stopped")
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Just stooped the ticker")
	fmt.Scanln()
}
