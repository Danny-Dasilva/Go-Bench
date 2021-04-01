package concurrent 
import (
    "net/http"
    "fmt"
    "time"
    "encoding/json"
    "sync"
)
type Comic struct {
    Num   int    `json:"num"`
    Link  string `json:"link"`
    Img   string `json:"img"`
    Title string `json:"title"`
}

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

func getComic(comicID int) (comic *Comic, err error) {
    url := fmt.Sprintf(baseXkcdURL, comicID)
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    err = json.NewDecoder(response.Body).Decode(&comic)
    if err != nil {
        return nil, err
    }

    return comic, nil
}

func main() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()

    comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    comicMap := make(map[int]*Comic, len(comicsNeeded))
    wg := sync.WaitGroup{}

    for _, id := range comicsNeeded {
        wg.Add(1)
        go func(id int) {
            comic, err := getComic(id)

            if err != nil {
                return
            }

            comicMap[id] = comic
            fmt.Printf("Fetched comic %d with title %v\n", id, comic.Title)
            wg.Done()
        }(id)
    }

    wg.Wait()
}
