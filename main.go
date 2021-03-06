package main

import (
	"log"
	"net"

	"net/http"

	"github.com/TykTechnologies/tyk-protobuf/bindings/go"
	"google.golang.org/grpc"
)

const (
	ListenAddress   = ":9111"
	ManifestAddress = ":8888"
)

func main() {
	lis, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening...", ListenAddress)
	s := grpc.NewServer()
	coprocess.RegisterDispatcherServer(s, &Dispatcher{})
	go s.Serve(lis)

	http.HandleFunc("/bundle.zip", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request for manifest")
		http.ServeFile(w, r, "bundle.zip")
	})

	http.ListenAndServe(ManifestAddress, nil)
}
