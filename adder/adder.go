package adder

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/apcera/nats"
)

// A request message that is sent to the Adder
type Request struct {
	X int64 `json:"x,omitempty"`
	Y int64 `json:"y,omitempty"`
}

// The response from the adders.
// Includes the answer and the UUID of the responder.
type Response struct {
	Ans int64  `json:"ans,omitempty"`
	Id  string `json:"id,omitempty"`
}

// The Adder
type Adder struct {
	// A unique Id that allows us to identify specific responders.
	Id string

	// The encoded nats connection
	nc *nats.EncodedConn
}

// This will create and run an Adder via a NATS subscription
// and callback, which happen in their own Go routine.
func NewAdder(subject string) *Adder {
	// Create the Adder
	a := &Adder{Id: genId(), nc: NatsConn()}

	// Create the subscriber for processing requests.
	a.nc.Subscribe(subject, a.processAddRequest)

	// Log that we are ready.
	fmt.Printf("Adder [%s] is ready\n", a.Id)

	return a
}

// processAddRequest will be responsible for adding the numbers if it
// can, and responding with its identity along with the answer itself.
func (a *Adder) processAddRequest(subject, reply string, r *Request) {
	resp := Response{Ans: (r.X + r.Y), Id: a.Id}
	a.nc.Publish(reply, &resp)
}

// Helper function to create an encoded connection to the messaging system.
func NatsConn() *nats.EncodedConn {
	// Connect to the messaging system
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Printf("Could not connect: %v\n", err)
		os.Exit(1)
	}

	// Create an encoded connection that sends payloads as JSON.
	nc, err := nats.NewEncodedConn(conn, "json")
	if err != nil {
		fmt.Printf("Could not create encoded connection: %v\n", err)
		os.Exit(1)
	}

	return nc
}

// A simple UUID random generator
func genId() string {
	u := make([]byte, 4)
	io.ReadFull(rand.Reader, u)
	return hex.EncodeToString(u)
}
