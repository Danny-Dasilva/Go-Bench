package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
    for j := range jobs {
        // fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        if id == 2 {
            time.Sleep(8 * time.Second)
        }
        l := fmt.Sprintf("worker", id, "finished job", j)
        results <- l
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan string, numJobs)

    for w := 1; w <= 10; w++ {
        go worker(w, jobs, results)
    }
    i := 1
    for {
        // for j := 1; j <= numJobs; j++ {
        i++
        jobs <- i
        
        // close(jobs)
    
        // for a := 1; a <= numJobs; a++ {
        r := <-results
        if r {
            fmt.Println(r)
        } else {
            continue
        }

    }
    
}