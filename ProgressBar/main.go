package main

import (
	"fmt"
	"sync"
	"time"
)

type Progress struct {
	TimeToComplete  time.Time
	CurrentProgress int
}

func progressCalculator(waitTime int, noOfDevices int, progress chan Progress, wg *sync.WaitGroup) {
	defer wg.Done()

	// ticker for estimated time updates
	ticker1 := time.NewTicker(500 * time.Millisecond)

	// ticker for progress updates (based on waitTime)
	ticker2 := time.NewTicker(time.Duration(waitTime) * time.Millisecond)

	count := 0

	for {
		select {

		// send estimated completion time
		case <-ticker1.C:
			progress <- Progress{
				TimeToComplete: time.Now().Add(10 * time.Second),
			}

		// send progress updates
		case <-ticker2.C:
			count += 10
			progress <- Progress{
				CurrentProgress: count,
			}

			// stop when progress reaches 100%
			if count >= 100 {
				close(progress)
				return
			}
		}
	}
}

func printProgress(progress chan Progress, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range progress {
		fmt.Println("Progress Update:", p)
	}
}

func main() {
	var wg sync.WaitGroup

	noOfDevices := 8
	waitTime := 20

	progressChan := make(chan Progress)

	wg.Add(2)

	go progressCalculator(waitTime, noOfDevices, progressChan, &wg)
	go printProgress(progressChan, &wg)

	wg.Wait()
}
