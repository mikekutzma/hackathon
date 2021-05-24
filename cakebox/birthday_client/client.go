// Client
package main

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/mikekutzma/hackathon/cakebox/cakebox"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultMonth = 4
	defaultDay   = 6
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCakeBoxClient(conn)

	// Parse input arguments
	month := defaultMonth
	day := defaultDay
	if len(os.Args) > 1 {
		dateParts := strings.Split(os.Args[1], "/")
		if len(dateParts) < 2 {
			log.Fatalf("Invalid date %s, should be in format m/d", os.Args[1])
		}
		month, _ = strconv.Atoi(dateParts[0])
		day, _ = strconv.Atoi(dateParts[1])
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create request object
	birthday := pb.Birthday{Month: int64(month), Day: int64(day)}

	// Perform request and get stream to start processing
	stream, err := c.UsersFromBirthday(ctx, &birthday)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	// Set a flag that just lets us know if we actually received anything
	wasEmpty := true
	for {
		// Pluck item off stream
		user, err := stream.Recv()

		// If error is end of stream, just means no more to process
		if err == io.EOF {
			break
		}

		wasEmpty = false

		if err != nil {
			log.Fatalf("%v.UsersFromBirthday(_) = _, %v", c, err)
		}

		log.Printf("User: %s", user)
	}
	// Drop a message to let user know nothing was returned
	if wasEmpty {
		log.Printf("No Users returned for %s", &birthday)
	}
}
