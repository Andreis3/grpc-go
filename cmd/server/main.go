package main

import (
	"database/sql"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"

	"github.com/andreis3/grpc-go/internal/database"
	"github.com/andreis3/grpc-go/internal/pb"
	"github.com/andreis3/grpc-go/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
