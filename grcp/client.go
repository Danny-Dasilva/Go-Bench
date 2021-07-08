package main

import (
	"context"
	"flag"
	
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	pb "github.com/Danny-Dasilva/gRPC-Tests/bidirectional/js-test/cycletlsproto"
)


var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)
func runCycleTLS(client pb.CycleStreamClient) {
	requests := []*pb.CycleTLSRequest{
		{RequestID: "1", Options: &pb.Options{URL: "https://www.google.com", Method: "GET", Headers: "", Body: "", Ja3: "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", Proxy: "", Cookies: ""}},
		{RequestID: "2", Options: &pb.Options{URL: "https://www.google.com", Method: "GET", Headers: "", Body: "", Ja3: "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", Proxy: "", Cookies: ""}},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Stream(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	n:= 0
	start := time.Now()
	defer func() {
		log.Println("Execution Time: ", time.Since(start))
	}()
	go func() {
		for n < 10000 {
			in, _ := stream.Recv()
			// if err == io.EOF {
			// 	// read done.
			// 	close(waitc)
			// 	return
			// }
			// if err != nil {
			// 	log.Fatalf("Failed to receive a note : %v", err)
			// }
			_= in
			if in != nil {

				// log.Printf("Got message %s at point(%d, %s)", in.RequestID, in.Status, in.Body)
				n++
				// log.Println(n)
				
				
			}
			
		}
		log.Println("done")
		close(waitc)
	}()
	// for _, request := range requests {
	for i := 1; i < 10002; i++ {
		if err := stream.Send(requests[0]); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}



func main() {
	flag.Parse()
	
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCycleStreamClient(conn)

	// RouteChat
	runCycleTLS(client)
}

