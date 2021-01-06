package main

import (
	"fmt"
	"time"
)

//var numCores = flag.Int("n", 2, "number of CPU cores to use")

//func main() {
//
//	//flag.Parse()
//	//runtime.GOMAXPROCS(*numCores)
//	fmt.Println(*numCores)
//
//	fmt.Println("In main()")
//	go longWait()
//	go shortWait()
//	fmt.Println("About to sleep in main()")
//	// sleep works with a Duration in nanoseconds (ns) !
//	time.Sleep(10 * 1e9)
//	fmt.Println("At the end of main()")
//}
//
//func longWait() {
//	fmt.Println("Beginning longWait()")
//	time.Sleep(5 * 1e9) // sleep for 5 seconds
//	fmt.Println("End of longWait()")
//}
//
//func shortWait() {
//	fmt.Println("Beginning shortWait()")
//	time.Sleep(2 * 1e9) // sleep for 2 seconds
//	fmt.Println("End of shortWait()")
//}

func main() {

	ch := make(chan string)

	sendData(ch)
	//go sendData(ch)
	go getData(ch)
	//go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}