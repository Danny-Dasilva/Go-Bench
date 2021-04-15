package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "runtime"
    "time"
    "io/ioutil"
    "strings"
)

var (
    reqs int
    max  int
)

func init() {
    flag.IntVar(&reqs, "reqs",  10000, "Total requests")
    flag.IntVar(&max, "concurrent", 100, "Maximum concurrent requests")
}

type Response struct {
    *http.Response
    err error
}

// Dispatcher
func dispatcher(reqChan chan *http.Request) {
    defer close(reqChan)
    for i := 0; i < reqs; i++ {
        req, err := http.NewRequest("GET", "http://localhost:8080", nil)
        if err != nil {
            log.Println(err)
        }
        reqChan <- req
    }
}

// Worker Pool
func workerPool(reqChan chan *http.Request, respChan chan string) {

    for i := 0; i < max; i++ {
        go worker(reqChan, respChan)
    }
}

// Worker
func worker(reqChan chan *http.Request, respChan chan string) {
    for req := range reqChan {
        // resp, err := http.Get("http://localhost:8081")
        client := &http.Client{}
        rr, err := http.NewRequest("GET", "http://localhost:8081", strings.NewReader(""))
        resp, err := client.Do(rr)

        // r := Response{resp, err}
        _ = req
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Print("bodyBytes" + err.Error())
        }
        r := string(bodyBytes)
       


        respChan <- r
    }
}

func main() {

    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()


    flag.Parse()
    runtime.GOMAXPROCS(runtime.NumCPU())
    
	// runtime.GOMAXPROCS(1)
	reqChan := make(chan *http.Request)
    respChan := make(chan string)
    
    // go dispatcher(reqChan)
    go workerPool(reqChan, respChan)
    
    go dispatcher(reqChan)




    var (
        conns int64

    )
   
    i := 0
    for conns < int64(reqs) {

        select {
        case message := <-respChan:
           i ++
           fmt.Println(i)
            _ = message
            conns++
        }
       
    }




    took := time.Since(start)
    ns := took.Nanoseconds()
    av := ns / conns
    average, err := time.ParseDuration(fmt.Sprintf("%d", av) + "ns")
    if err != nil {
        log.Println(err)
    }
    fmt.Printf("Connections:\t%d\nConcurrent:\t%d\nTotal time:\t%s\nAverage time:\t%s\n", conns, max, took, average)
}