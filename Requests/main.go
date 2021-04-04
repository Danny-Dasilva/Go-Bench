package main 
import (
	//"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"time"
    "sync"
    "net/http"
	"github.com/valyala/fasthttp"
	//"github.com/valyala/fastjson"
)
type Comic struct {
    Num   int    `json:"num"`
    Link  string `json:"link"`
    Img   string `json:"img"`
    Title string `json:"title"`
}

const baseURL = "http://localhost:8080"
// const baseURL = "https://www.digitalocean.com/"
func getFast(comicID int) (comic string, err error) {
    _, body, err  := fasthttp.Get(nil, baseURL)
    if err != nil {
        return "aa", err
    }

    bodyBytes, err := ioutil.ReadAll(bytes.NewBuffer(body))

    comic = string(bodyBytes) 
    return comic, nil
}

func get(comicID int) (comic string, err error) {
    response, err := http.Get(baseURL)
    if err != nil {
        return "aa", err
    }

    bodyBytes, err := ioutil.ReadAll(response.Body)

    comic = string(bodyBytes) 
    return comic, nil
}


func fast_concurrent() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

    //comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    wg := sync.WaitGroup{}
    
    for i := 1; i <= 10000; i++ {
        wg.Add(1)
        go func(i int) {
            _, err := getFast(i)

            if err != nil {
                return
            }

            fmt.Println("running", i, "\n")
            wg.Done()
        }(i)
    }

    wg.Wait()
}

func concurrent() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

    //comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    wg := sync.WaitGroup{}
    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go func(i int) {
            _, err := get(i)

            if err != nil {
                return
            }

            fmt.Println("running", i, "\n")
            wg.Done()
        }(i)
    }

    wg.Wait()
}


func nonconcurrent() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

  

    for i := 1; i <= 100; i++ {
        _, err := get(i)
        if err != nil {
            continue
        }

        fmt.Println("running", i, "\n")
    }
}

func fast_nonconcurrent() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

  

    for i := 1; i <= 100; i++ {
        _, err := getFast(i)
        if err != nil {
            continue
        }

        fmt.Println("running", i, "\n")
    }
}

func main() {
    concurrent()
}