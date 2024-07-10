package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/nezo32/sudoku/iam/generated/protos/user"
	"github.com/nezo32/sudoku/iam/services/rpc/user_service"
)

func main() {
	godotenv.Load()

	fmt.Println("Connecting to database...")
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Addr:     os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	})
	defer db.Close()

	ctx := context.Background()

	fmt.Println("Check database health...")
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("Start listening on %s...\n", os.Getenv("ENDPOINT_RPC_PORT"))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("ENDPOINT_RPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	grpc_server := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc_server, user_service.CreateUserSerivceServer(e, db))
	reflection.Register(grpc_server)

	fmt.Println("Starting http/rpc server...")
	go grpc_server.Serve(lis)
	e.Logger.Fatal(e.Start(":" + os.Getenv("ENDPOINT_PORT")))
}
