// package main

// import (
// 	//"encoding/json"
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"sync"
// 	"time"

// 	"github.com/valyala/fasthttp"
// 	//"github.com/valyala/fastjson"
// )

// //type Comic struct {
// //Num   int    `json:"num"`
// //Link  string `json:"link"`
// //Img   string `json:"img"`
// //Title string `json:"title"`
// //}

// const baseURL = "http://localhost:8080"

// func getComic(comicID int) (comic string, err error) {
//     _, body, err  := fasthttp.Get(nil, baseURL)
//     if err != nil {
//         return "aa", err
//     }
//     //t := json.Unmarshal(body, &dat) 

//     bodyBytes, err := ioutil.ReadAll(bytes.NewBuffer(body))

//     //ioutil.ReadAll(resp.Body)
//     //err = json.NewDecoder(decoded).Decode(&comic)
//     //if err != nil {
//         //return nil, err
//     //}

//     comic = string(bodyBytes) 
//     return comic, nil
// }

// func main() {
//     start := time.Now()
//     defer func() {
//         fmt.Println("Execution Time: ", time.Since(start))
//     }()

//     //comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//     wg := sync.WaitGroup{}

//     for i := 1; i <= 1000; i++ {
//         wg.Add(1)
//         go func(i int) {
//             _, err := getComic(i)

//             if err != nil {
//                 return
//             }

//             fmt.Println("running", i, "\n")
//             //fmt.Printf("Fetched comic %d with title %v\n", i, comic)
//             wg.Done()
//         }(i)
//     }

//     wg.Wait()
// }
