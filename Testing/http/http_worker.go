
package main
import (
	"fmt"
	"time"
    "runtime"
    "net/http"
    "io/ioutil"

	// "os"
	// "strconv"
	
)


type Response struct {
    *http.Response
    err error
}


const baseURL = "http://localhost:8080"





func process(job int, i int, link chan<- string) {

    req, err := http.NewRequest("GET", "http://localhost:8080", nil)
    if err != nil {
        fmt.Print( "req" + err.Error())

    }
    resp, err := http.DefaultTransport.RoundTrip(req)
    
    defer resp.Body.Close()
        
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Print("bodyBytes" + err.Error())
    }
    r := string(bodyBytes)

	link <- r
}
	
func worker(jobChan <-chan int, i int, link chan<- string) {
	for job := range jobChan {
		process(job, i, link)
	}
}




func main() {

    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

	workerCount := 200
    runtime.GOMAXPROCS(runtime.NumCPU())

    requests := 10000

    
	// make a channel with a capacity of 100.
	jobChan := make(chan int, 100) // Or jobChan := make(chan int)
	// done := make(chan bool)
	link := make(chan string)
	// start the worker
	for i:=0; i<workerCount; i++ {
		go worker(jobChan, i, link)
	}
	


	ch := make(chan int)
    go func(ch chan int) {
   
        for i := 0; i < requests; i++ {
            
            ch <- i
        }
        
        // Simulating stdin
        close(ch)
    }(ch)


    for stdin := range ch {
        jobChan <- stdin
    }
    var conns int
    for conns < requests {

        select {
        case message := <-link:
            fmt.Println(message)
            conns++
        }
        // fmt.Println(conns)
    }

	
	

	fmt.Println("done")
}