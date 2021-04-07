
package main
import (
	"fmt"
	"time"
	// "os"
	// "strconv"
	"bufio"
	"os"
)







func process(job string, i int, link chan<- string) {

	if job == "2\n" {
		time.Sleep(2 *  time.Second)
		
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
	// done := make(chan bool)
	link := make(chan string)
	// start the worker
	for i:=0; i<workerCount; i++ {
		go worker(jobChan, i, link)
	}
	


	ch := make(chan string)
    go func(ch chan string) {
        // Uncomment this block to actually read from stdin
        reader := bufio.NewReader(os.Stdin)
        for {
            s, err := reader.ReadString('\n')
            if err != nil { // Maybe log non io.EOF errors, if you want
                close(ch)
                return
            }
            ch <- s
        }
        
        // Simulating stdin
        ch <- "A line of text"
        close(ch)
    }(ch)




	for {



		select {
        case stdin, ok := <-ch:
            if !ok {
                break 
            } else {
				jobChan <- stdin
            }
        case <-time.After(100 * time.Millisecond):
            // Do something when there is nothing read from stdin
        }

		
		select {
        case message := <-link:
            fmt.Println(message)
        default:

        }

		

	}
	
	

	fmt.Scanln()
}