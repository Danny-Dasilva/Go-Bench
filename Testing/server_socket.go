package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
	// "sync"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
type myTLSRequest struct {
	RequestID string `json:"requestId"`
	Options   struct {
		URL     string            `json:"url"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
		Ja3     string            `json:"ja3"`
		UserAgent     string       `json:"userAgent"`
		ID     int            		`json:"id"`
		Proxy   string            `json:"proxy"`
	} `json:"options"`
}

type options struct {
	ID     int            		`json:"id"`
}
type response struct {
	Status  int
	Body    string
	Headers map[string]string
}


type myTLSResponse struct {
	RequestID string
	Response  response
}


// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	ch := make(chan *myTLSResponse)
    go func(ch chan  *myTLSResponse) {
        // Uncomment this block to actually read from stdin
        for {
            mytlsresponse := new(myTLSResponse)
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			e := json.Unmarshal(message, &mytlsresponse)
			if e != nil {
				log.Print(e)
			}


			
            ch <- mytlsresponse
        }
        
       
    }(ch)
	r := 0 
	i := 0
	start := time.Now()
	defer func() {
		log.Println("Execution Time: ", time.Since(start))
	}()
	go func() {
		for i < 10000 {
			mytlsrequest := new(myTLSRequest)
			mytlsrequest.RequestID = string('t')
			mytlsrequest.Options.URL = "http://localhost:8081"
			mytlsrequest.Options.Method = "GET"
			// mytlsrequest.Options.Headers = map[string]string{
			// 							"Access-Control-Allow-Credentials": "Bearer someexampletoken",

			// 								}
			
			mytlsrequest.Options.Ja3 = "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0"

			mytlsrequest.Options.UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0"
			data, err := json.Marshal(mytlsrequest)
			if err != nil {
				log.Print("Request_Id_On_The_Left" )
				
			}

			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Print("Request_Id_On_The_Left" )
			}
		}
	}()
	for r < 10000{

	
			select {
			case message := <-ch:
				r ++
				// log.Println(message)
				_ = message
				
			default:
	
			}


			
		// }(i)

	}
	log.Println("done")
	// wg.Wait()
	
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9112", nil))
}
