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
	Options   options
}
type options struct {
	ID     int            		`json:"id"`
}
type response struct {
	Status  int
	Body    string
}

type myTLSResponse struct {
	RequestID string
	Response  response
}


// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	// wg := sync.WaitGroup{}
	for i:=0; i<10; i++ {
		// wg.Add(1)
		// go func(i int) {
			// read in a message
			
			options := options{i}
			write := myTLSRequest{"ext", options}

			// print out that message for clarity
			// log.Println(string(p))

			data, err := json.Marshal(write)
			if err != nil {
				log.Print("Request_Id_On_The_Left" )
				
			}

			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Print("Request_Id_On_The_Left" )
				
			}


			// mytlsresponse := new(myTLSResponse)
			// _, message, err := conn.ReadMessage()
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }
			// e := json.Unmarshal(message, &mytlsresponse)
			// if e != nil {
			// 	log.Print(e)
			// }


			// fmt.Println(mytlsresponse)
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
