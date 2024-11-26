// Go program to illustrate 
// how to create a channel 
package main 

import ("fmt"
"time"
)

func main() { 

	// Creating a channel using make() function 
	channel := make(chan int)

	// Call a function using go routines
	go simpleSleep(channel)

	// Wait for the function to send back data
	fmt.Println(<- channel)

	close(channel)
} 

func simpleSleep(channel chan int) {
	
	// Sleep for 5 seconds
	time.Sleep(time.Second * 5)

	// Send back a value
	channel <- 5
}


