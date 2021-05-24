// Client
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/mikekutzma/hackathon/cakebox/cakebox"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	defaultPort   = getEnv("PORT", "50051")
	grpcEndpoint  = flag.String("grpc-endpoint", fmt.Sprintf("localhost:%s", defaultPort), "The gRPC Endpoint of the Server")
	inputBirthday = flag.String("date", "", "Birthday to get users, formatted as m/d")
	noTLS         = flag.Bool("no-tls", false, "Flag to disable TLS (use this for requests to localhost container)")
)

// Helper function to get env variable with a default
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	flag.Parse()
	if *inputBirthday == "" {
		log.Fatal("Must specify --date input arg of format m/d")
	}

	// Set up connection options
	opts := []grpc.DialOption{
		grpc.WithBlock(),
	}

	if *noTLS {
		opts = append(opts, grpc.WithInsecure())
	} else {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*grpcEndpoint, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCakeBoxClient(conn)

	// Parse input arguments
	dateParts := strings.Split(*inputBirthday, "/")
	if len(dateParts) < 2 {
		log.Fatalf("Invalid date %s, should be in format m/d", *inputBirthday)
	}
	month, _ := strconv.Atoi(dateParts[0])
	day, _ := strconv.Atoi(dateParts[1])

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
