package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brotherlogic/goserver"
	"github.com/ianr0bkny/go-sonos/ssdp"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/sonosrpc/proto"
)

// Server the configuration for the syncer
type Server struct {
	*goserver.GoServer
	mgr deviceManager
}

// HTTPRetriever pulls http pages
type HTTPRetriever struct{}

// Does a web retrieve
func (r *HTTPRetriever) retrieve(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// DoRegister does RPC registration
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterSonosServiceServer(server, &s)
}

// ReportHealth reports health of the server
func (s Server) ReportHealth() bool {
	return true
}

// InitServer builds an initial server
func InitServer() Server {
	server := Server{&goserver.GoServer{}, ssdp.MakeManager()}
	server.Register = server

	return server
}

func main() {
	server := InitServer()

	server.PrepServer()
	server.RegisterServer("sonosbridge", false)
	server.Serve()
}
