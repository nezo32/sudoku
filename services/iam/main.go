package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbRoles "github.com/nezo32/sudoku/iam/generated/protos/roles"
	pb "github.com/nezo32/sudoku/iam/generated/protos/user"
	"github.com/nezo32/sudoku/iam/security"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/auth"
	"github.com/nezo32/sudoku/iam/services/http/roles"
	"github.com/nezo32/sudoku/iam/services/rpc/roles_service"
	"github.com/nezo32/sudoku/iam/services/rpc/user_service"
)

func main() {
	godotenv.Load()

	argonParams := security.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  2,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	hasher := security.HashPasswordGenerator(argonParams)

	jwtGenerator := security.JWTFactory()

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
		log.Error(err)
		log.Error(err)
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Check database health...")
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("Start listening on %s...\n", os.Getenv("ENDPOINT_RPC_PORT"))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("ENDPOINT_RPC_PORT")))
	if err != nil {
		log.Error(err)
		log.Error(err)
		log.Fatalf("failed to listen: %v", err)
		return
	}

	e := echo.New()
	e.Use(middleware.Recover())

	serviceCtx := &services.ServiceContext{
		Database:       conn,
		Echo:           e,
		Context:        ctx,
		JWTGenerator:   jwtGenerator,
		PasswordEncode: hasher,
	}

	auth.DefineAuthEndpoints(serviceCtx)
	roles.DefineRolesEndpoints(serviceCtx)

	grpc_server := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc_server, user_service.CreateUserSerivceServer(serviceCtx))
	pbRoles.RegisterRolesServiceServer(grpc_server, roles_service.CreateRolesServiceServer(serviceCtx))
	reflection.Register(grpc_server)

	fmt.Println("Starting http/rpc server...")
	go grpc_server.Serve(lis)
	e.Logger.Fatal(e.Start(":" + os.Getenv("ENDPOINT_PORT")))
}
