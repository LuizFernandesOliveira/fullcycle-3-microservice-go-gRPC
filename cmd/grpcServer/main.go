package main

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-gRPC/internal/database"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-gRPC/internal/pb"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
