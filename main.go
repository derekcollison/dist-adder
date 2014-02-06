//
// Sample to show the power of message based communication patterns in distributed systems.
//

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	gnatsd "github.com/apcera/gnatsd/server"
	"github.com/apcera/nats"
	"github.com/derekcollison/dist-adder/adder"
)

const (
	DefaultResponders = 10
	DefaultRequests   = 10

	ReqSub = "gophers.add.request"
)

func main() {

	// Flags
	var numResponders int
	var numRequests int

	flag.IntVar(&numResponders, "numResponders", DefaultResponders, "Responders to spin up.")
	flag.IntVar(&numRequests, "numRequests", DefaultRequests, "Requests to send.")
	flag.Parse()

	// Start the NATS server.
	startNatsServer()

	// Spin up the appropriate number of responders.
	fmt.Printf("\nSpinning up %d responders.\n\n", numResponders)
	for i := 0; i < numResponders; i++ {
		adder.NewAdder(ReqSub)
	}

	// Grab a client connection for sending requests.
	nc := adder.NatsConn()

	// Time to wait for a response before timing out.
	ttl := 10 * time.Millisecond

	var req *adder.Request
	var resp adder.Response

	// Send some requests.
	fmt.Printf("\nSending %d requests.\n\n", numRequests)
	for i := 0; i < numRequests; i++ {
		req = &adder.Request{X: rand.Int63() % 100, Y: rand.Int63() % 100}
		nc.Request(ReqSub, req, &resp, ttl)
		fmt.Printf("Request: %+v\tResponse: %+v\n", *req, resp)
	}

	fmt.Printf("\nFinished\n\n")
}

// Spin up the gnatsd server. We run the server in a Go routine
// here for simple startup and cleanup. Normally this is run as
// a separate process or cluster of processes.
func startNatsServer() {
	o := &gnatsd.Options{Port: nats.DefaultPort, NoLog: true}
	natsServer := gnatsd.New(o)
	go natsServer.Start()
}
