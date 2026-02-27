package main

import (
	"fmt"
	"time"
)

func player(name string, table chan string) {
	for {
		msg := <-table
		fmt.Println(name, msg)
		time.Sleep(500 * time.Millisecond)
		table <- "ball"
	}
}

func main() {
	table := make(chan string)

	go player("Ping", table)
	go player("Pong", table)

	// Start game
	table <- "ball"

	// Timeout after 5 seconds
	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("\nGame Over (Timeout)")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
