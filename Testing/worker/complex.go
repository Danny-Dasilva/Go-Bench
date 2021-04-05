package main

import (
	"fmt"
)


// Producers struct 
// payload to be ferried to the consumers and 
//ID of the producer that delivers the payload
type Producers struct {
	payloadChan	chan string  
	ID			int
}

// The job to be done
var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

// Assign jobs to the consumers with the help of the execute function
// workerPool ( a channel of type *Producers) is the only link betwen our producers to the consumers
// When a job is received, send a producer to the pool and load the producers payload with the job

func produce(jobChan <-chan string, producer *Producers, workerPool chan *Producers){

	for {

		select {
			case job := <- jobChan: {
				if len(job) > 0 {  

					workerPool<- producer
					producer.payloadChan <-job

					fmt.Printf("Producer %d ......> %s\n\n", producer.ID, job)
				}
			}
		}
	}
}

// listen for any produced job (from producer)
func consume(id int, workerPool chan *Producers){
 
	// an infinite loop listening for each consumer
	for { 

		worker := <- workerPool

		if job, ok := <- worker.payloadChan ; ok { 
			fmt.Printf("Consumer %d ---> %s\n\n", id, job) 
			
		}
	}  
}

// send jobs to jobChan 
//which will be recieved by producer function
func execute(jobChan chan<- string, jobs []string, allDone chan<- bool){

	for _, job := range jobs {
		jobChan<- job
	}

	defer close(jobChan)
	allDone<- true
}

func main() {

	// a slice of jobs. contains number of simultaneous/concurent routines
	// you wish to run at a time
	var workers [ ]*Producers 

	jobChan := make(chan string)
	workerPool := make(chan *Producers)
	allDone := make(chan bool)
	
	// you can basically tweak this number as you wish 
	producers := 10
	consumers := 10 
	
	// launch producers
	for i := 0; i <= producers; i ++ { 
		// append a new worker
		workers = append(workers, 
			&Producers{
				payloadChan : make(chan string),
				ID: i,
			},
		)
		
		go produce(jobChan, workers[i], workerPool)
	}
	
	// assign jobs to the ready workers 
	go execute(jobChan, messages, allDone)

	// launch consumers that will listen from workerPool
	for c := 0; c <= consumers; c++ { 
		go consume(c, workerPool)
	}
	
	<-allDone
}