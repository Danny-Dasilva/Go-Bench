// // concurrent.go
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// 	"io/ioutil"
// )
// const baseURL = "http://localhost:8080"
// func MakeRequest(url string, ch chan<-string, num int) {
// 	start := time.Now()
// 	resp, _ := http.Get(url)  
// 	secs := time.Since(start).Seconds()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()

// 	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)

// }

// func main() {

// 	start := time.Now()
//     defer func() {
//         fmt.Println("Execution Time: ", time.Since(start))
//     }()

// 	ch := make(chan string)
// 	for i := 1; i <= 10000; i++ {
// 		go MakeRequest(baseURL, ch, i)
// 		fmt.Println(i)
// 	}  
// 	for i := 1; i <= 10000; i++ {
// 		fmt.Println(<-ch)
// 	}
	
 	
// }