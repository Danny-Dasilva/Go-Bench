
// package main
// import (
// 	"fmt"
// 	"time"
//     "runtime"
//     "net/http"
//     "io/ioutil"

// 	// "os"
// 	// "strconv"
	
// )



// const baseURL = "http://localhost:8080"




// // func process(job string, i int, link chan<- string) {

// // 	if job == "2\n" {
// // 		time.Sleep(2 *  time.Second)
		
// // 	}
// // 	r := fmt.Sprintf("%v by Job%v\n", job, i)
// // 	link <- r
// // }
	
// // func worker(jobChan <-chan string, i int, link chan<- string) {
// // 	for job := range jobChan {
// // 		process(job, i, link)
// // 	}
// // }

// func process(job int, i int, link chan<- string) {

// 	response, err := http.Get(baseURL)
//     if err != nil {
//         fmt.Println(err)
//     }

//     bodyBytes, err := ioutil.ReadAll(response.Body)

//     r := string(bodyBytes) 
// 	link <- r
// }
	
// func worker(jobChan <-chan int, i int, link chan<- string) {
// 	for job := range jobChan {
// 		process(job, i, link)
// 	}
// }




// func main() {

//     start := time.Now()
//     defer func() {
//         fmt.Println("Execution Time: ", time.Since(start))
//     }()

// 	workerCount := 100
//     runtime.GOMAXPROCS(runtime.NumCPU())

//     requests := 100

    
// 	// make a channel with a capacity of 100.
// 	jobChan := make(chan int, 100) // Or jobChan := make(chan int)
// 	// done := make(chan bool)
// 	link := make(chan string)
// 	// start the worker
// 	for i:=0; i<workerCount; i++ {
// 		go worker(jobChan, i, link)
// 	}
	


// 	ch := make(chan int)
//     go func(ch chan int) {
   
//         for i := 0; i < requests; i++ {
            
//             ch <- i
//         }
        
//         // Simulating stdin
//         close(ch)
//     }(ch)


//     for stdin := range ch {
//         jobChan <- stdin
//     }
//     var conns int
//     for conns < requests {

//         select {
//         case message := <-link:
//             fmt.Println(message)
//             conns++
//         }
//         fmt.Println(conns)
//     }

	
	

// 	fmt.Println("done")
// }