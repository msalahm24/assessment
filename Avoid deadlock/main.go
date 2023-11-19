package main

import (
	"fmt"
	"time"
)

func main() {
	RunEightReadingTwoWriting(16,16)
}

func RunEightReadingTwoWriting(readers int,writers int) {
	buffer := make(chan int, 1)
	readDone := make(chan bool)
	writeDone := make(chan bool)

	for i := 0; i < writers; i++ {
		go func(id int) {
			for {
				buffer <- id
				fmt.Printf("Writer %d wrote to the buffer\n", id)
				time.Sleep(time.Second)
				writeDone <- true
			}
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func(id int) {
			for {
				value := <-buffer
				fmt.Printf("Reader %d read from the buffer: %d\n", id, value)
				time.Sleep(time.Second)
				readDone <- true
			}
		}(i)
	}

	
	for {
		select {
		case <-readDone:
			fmt.Println("A reader completed")
		case <-writeDone:
			fmt.Println("A writer completed")
		}
	}
}


