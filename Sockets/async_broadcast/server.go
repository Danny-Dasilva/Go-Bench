package main

import (
	"fmt"
	"time"
	"github.com/gorilla/websocket"
	"net/http"
)


func client(clientName string, clientChan chan string) {
    for {
        text, _ := <-clientChan
        fmt.Printf("%s: %s\n", clientName, text)
    }
}




func server(serverChan chan chan string) {
    var clients []chan string
    for {
        select {
        case client, _ := <-serverChan:
            clients = append(clients, client)
            // Broadcast the number of clients to all clients:
            for _, c := range clients {
                c <- fmt.Sprintf("%d client(s) connected.", len(clients))
            }
        }
    }
}

func uptimeServer(serverChan chan chan string) {
    var clients []chan string
    uptimeChan := make(chan int, 1)
    // This goroutine will count our uptime in the background, and write
    // updates to uptimeChan:
    go func (target chan int) {
        i := 0
        for {
            time.Sleep(time.Second)
            i++
            target <- i
        }
    }(uptimeChan)
    // And now we listen to new clients and new uptime messages:
    for {
        select {
        case client, _ := <-serverChan:
            clients = append(clients, client)
        case uptime, _ := <-uptimeChan:
            // Send the uptime to all connected clients:
            for _, c := range clients {
                c <- fmt.Sprintf("%d seconds uptime", uptime)
            }
        }
    }
}

var upgrader = websocket.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r * http.Request) bool {
        return true // Disable CORS for testing
    },
}



func main() {
	serverChan := make(chan chan string, 4)
	go uptimeServer(serverChan)


	// Define a HTTP handler function for the /status endpoint, that can receive
	// WebSocket-connections only... so note that browsing it with your browser will fail.
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade this HTTP connection to a WS connection:
		ws, _ := upgrader.Upgrade(w, r, nil)
		// And register a client for this connection with the uptimeServer:
		client := make(chan string, 1)
		serverChan <- client
		// And now check for uptimes written to the client indefinitely.
		// Yes, we are lacking proper error and disconnect checking here, too:
		for {
			select {
			case text, _ := <-client:
				writer, _ := ws.NextWriter(websocket.TextMessage)
				writer.Write([]byte(text))
				writer.Close()
			}
		}
	})


	http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second)


}