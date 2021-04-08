package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

	i := 0
	// wg := sync.WaitGroup{}
	for {
		// wg.Add(1)
		// go func(i int) {
			// read in a message
			if i < 10 {
				mytlsrequest := new(myTLSRequest)
				mytlsrequest.RequestID = string('t')
				mytlsrequest.Options.URL = "http://localhost:8080"
				mytlsrequest.Options.Method = "GET"
				mytlsrequest.Options.Headers = map[string]string{
											"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",

												}
				
				mytlsrequest.Options.Ja3 = "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0"


				data, err := json.Marshal(mytlsrequest)
				if err != nil {
					log.Print("Request_Id_On_The_Left" )
					
				}

				err = conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Print("Request_Id_On_The_Left" )
				}
			}

			i++

	
			select {
			case message := <-ch:
				

				fmt.Println(message.RequestID)
			default:
	
			}


			
		// }(i)

	}
	// wg.Wait()
	
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	
	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
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
