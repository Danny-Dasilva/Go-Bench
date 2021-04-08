package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "runtime"
    "time"
    "io/ioutil"
)

var (
    reqs int
    max  int
)

func init() {
    flag.IntVar(&reqs, "reqs",  100, "Total requests")
    flag.IntVar(&max, "concurrent", 10000, "Maximum concurrent requests")
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
        resp, err := http.DefaultTransport.RoundTrip(req)
        // r := Response{resp, err}

        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Print("bodyBytes" + err.Error())
        }
        r := string(bodyBytes)



        respChan <- r
    }
}

// Consumer
func consumer(respChan chan Response) (int64, int64) {
    var (
        conns int64
        size  int64
    )
    for conns < int64(reqs) {
        select {
        case r, ok := <-respChan:
            if ok {
                
                if r.err != nil {
                    log.Println(r.err)
                } else {
                    size += r.ContentLength
                    if err := r.Body.Close(); err != nil {
                        log.Println(r.err)
                    }
                }
                conns++
            }
        }
    }
    return conns, size
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
    
    
    defer close(reqChan)

    ch := make(chan *http.Request)
    go func(ch chan *http.Request) {

        for i := 0; i < reqs; i++ {
            req, err := http.NewRequest("GET", "http://localhost:8080", nil)
            if err != nil {
                log.Println(err)
            }
            reqChan <- req
        }
    }(ch)


    // for stdin := range ch {
    //     fmt.Println("ahh")
    //     reqChan <- stdin
    // }


    var (
        conns int64
        size  int64
    )
   

    for conns < int64(reqs) {

        select {
        case message := <-respChan:
            // fmt.Println(message)
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
    fmt.Printf("Connections:\t%d\nConcurrent:\t%d\nTotal size:\t%d bytes\nTotal time:\t%s\nAverage time:\t%s\n", conns, max, size, took, average)
}