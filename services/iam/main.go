package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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

	ctx := context.Background()

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	conn, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	/* fmt.Println("Check database health...")
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	} */

	fmt.Printf("Start listening on %s...\n", os.Getenv("ENDPOINT_RPC_PORT"))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("ENDPOINT_RPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	grpc_server := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc_server, user_service.CreateUserSerivceServer(e, conn, ctx))
	reflection.Register(grpc_server)

	fmt.Println("Starting http/rpc server...")
	go grpc_server.Serve(lis)
	e.Logger.Fatal(e.Start(":" + os.Getenv("ENDPOINT_PORT")))
}
