
package main
import (
	"fmt"
	"time"
	// "os"
	"strconv"
)




func consumer(link <-chan string, done chan<- bool) {
	for b := range link {
		fmt.Println(b)
	}
	done <- true
}




func process(job string, i int, link chan<- string) {

	// if job == 5 {
	// 	timer := time.NewTimer(2 * time.Second)
	// 	<-timer.C
	// } else {
	// 	timer := time.NewTimer(time.Duration(job) * time.Millisecond)
	// 	<-timer.C
	// }
	if job == "2" {
		time.Sleep(4 *  time.Second)
		
	}
	
	r := fmt.Sprintf("%v by Job%v\n", job, i)
	link <- r
}
	
func worker(jobChan <-chan string, i int, link chan<- string) {
	for job := range jobChan {
		process(job, i, link)
	}
}

func main() {
	workerCount := 2
	// make a channel with a capacity of 100.
	jobChan := make(chan string, 100) // Or jobChan := make(chan int)
	done := make(chan bool)
	link := make(chan string)
	// start the worker
	for i:=0; i<workerCount; i++ {
		go worker(jobChan, i, link)
	}
	
	for i:=0; i<10; i++ {
		s := strconv.Itoa(i)
		jobChan <- s
	}
	

	go consumer(link, done)
	

	fmt.Scanln()
}