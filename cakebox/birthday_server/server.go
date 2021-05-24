// Server
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/mikekutzma/hackathon/cakebox/cakebox"
	"google.golang.org/grpc"

	_ "github.com/mattn/go-sqlite3"
)

// Create global db variable
var (
	port = getEnv("PORT", "50051")
	db   *sql.DB
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCakeBoxServer
}

type comparableBirthday struct {
	Month int
	Day   int
}

type comparableUser struct {
	Name string
}

// Helper function to get env variable with a default
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// SayHello implements helloworld.GreeterServer
func (s *server) UsersFromBirthday(in *pb.Birthday, stream pb.CakeBox_UsersFromBirthdayServer) error {
	log.Printf("Received: %s", in)
	query := `
		select
			name
		from
			birthdays
		where 1=1
			and birthMonth = ?
			and birthDay = ?
	`
	// Execute query, with placeholders
	rows, err := db.Query(query, in.GetMonth(), in.GetDay())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// Create placeholder for row values
		var name string
		// Assign values from query to placeholder
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		// Create new message object from value
		user := &pb.User{Name: name}
		// Send message on stream
		if err := stream.Send(user); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcEndpoint := fmt.Sprintf(":%s", port)
	log.Printf("gRPC endpoint [%s]", grpcEndpoint)
	// Listen on port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Register server
	s := grpc.NewServer()
	pb.RegisterCakeBoxServer(s, &server{})

	// Create database connection
	db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
