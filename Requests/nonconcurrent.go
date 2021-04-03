package main 
import (
	//"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

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

func getComic(comicID int) (comic string, err error) {
    _, body, err  := fasthttp.Get(nil, baseURL)
    if err != nil {
        return "aa", err
    }

    bodyBytes, err := ioutil.ReadAll(bytes.NewBuffer(body))

    comic = string(bodyBytes) 
    return comic, nil
}

func main() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

  

    for i := 1; i <= 10000; i++ {
        _, err := getComic(i)
        if err != nil {
      continue
        }

        fmt.Println("running", i, "\n")
    }
}
