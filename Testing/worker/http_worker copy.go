package main 

import (
	"fmt"
	// "bytes"
	"io/ioutil"
	"time"
    "net/http"

)
const baseURL = "http://localhost:8080"


func process(job string, i int, link chan<- string) {

	
	response, err := http.Get(baseURL)
    if err != nil {
        fmt.Println(err)
    }

    bodyBytes, err := ioutil.ReadAll(response.Body)

    r := string(bodyBytes) 

	link <- r
}

func worker(jobChan <-chan string, i int, link chan<- string) {
	for job := range jobChan {
		fmt.Println(i)
		process(job, i, link)
	}

	close(link)
}

func main() {


	start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()



	workerCount := 2
	// make a channel with a capacity of 100.
	jobChan := make(chan string, 100) // Or jobChan := make(chan int)
	// done := make(chan bool)
	link := make(chan string)
	// start the worker
	for i:=0; i<workerCount; i++ {
		go worker(jobChan, i, link)
	}




	ch := make(chan string)
	go func(ch chan string) {
        
        
		// for i := 1; i <= 10; i++ {
		// 	// enqueue a job
		// 	ch <- string(i)
		// }
		
        // Simulating stdin
        ch <- "A line of text"
		ch <- "done"
        close(ch)
    }(ch)

	
	
		



	for {



		select {
        case stdin, ok := <-ch:
            if !ok {
				ch = nil
                break 
            } else {
				jobChan <- stdin
            }
        default:
            // Do something when there is nothing read from stdin
        }

		
		select {
        case message, ok := <-link:
			if !ok {
				fmt.Println("yeet")
				link = nil
			} else if message == "done" {
				fmt.Println("yeet")
			}

            fmt.Println(message)
        default:
			

        }
		if ch == nil && link == nil{
			break
		}
		
		

	}


	fmt.Scanln()
}