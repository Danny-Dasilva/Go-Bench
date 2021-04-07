package main

import (
    "fmt"
    "net"
    "sync"
    "log"
)

func main() {
    fmt.Println("aaa")
    var wg sync.WaitGroup
    wg.Add(2)

    server := StartShard()
    fmt.Println(server)
    client := StartClient()
    fmt.Println(client, server)
    go func() {
        ShardToClient(client, server)
        wg.Done()
    }()

    go func() {
        ClientToShard(client, server)
        wg.Done()
    }()

    wg.Wait()
    fmt.Println("yteeyt")

}

func StartClient() (net.Conn){

    

    return conn
}

func StartShard() (net.Conn){


    return conn
}


func ShardToClient( client net.Conn, server net.Conn  ){

    var buf = make([]byte, 1024)
    var bufRead int
    var err error

    for {
        bufRead, err = client.Read(buf)
        checkError(err)

        if bufRead > 0 {
            if buf[0] == 0x8C{
                fmt.Println("AVOIDING REDIRECTION ...")
                buf[1] = 127;
                buf[2] = 0;
                buf[3] = 0;
                buf[4] = 1;
            }
            fmt.Printf("ShardToClient(%d): %X\n", bufRead, buf[0:bufRead])
            bufRead, _ = server.Write(buf[0:bufRead])

        }
    }
}
func ClientToShard( client net.Conn, server net.Conn  ){
    defer client.Close() // close network connections on return from this function
    defer server.Close()

    var buf = make([]byte, 1024)
    var bufRead int
    var err error

    for {
        bufRead, err = server.Read(buf)
        if err != nil {
            return
        }

        if bufRead > 0 {
            fmt.Printf("ClientToShard(%d): %X\n", bufRead, buf[0:bufRead])
            bufRead, err = client.Write(buf[0:bufRead])
            if err != nil {
                return
            }
        }
    }

}

func checkError(err error) {
    if err != nil {
        log.Fatal("fatal: %s", err)
    }
}