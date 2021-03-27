package main

import (
	"os"
	"net"
	"log"
	"context"
	//"errors"
	
	pb "github.com/hitesh-sureify/grpc-demo/proto"
	"github.com/hitesh-sureify/grpc-demo/db"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedEmployeeServiceServer
}

func main() {
	
	listen, err := net.Listen("tcp", os.Getenv("GRPC_SRV_ADDR"))
	if err != nil{
		log.Fatalf("Could not listen on port : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

	log.Printf("Hosting server on : %s", listen.Addr().String())

}



func (s * server) CreateEmployee(ctx context.Context, emp *pb.Employee) (*pb.ID, error){

	if id, err := db.Insert(ctx,emp); err != nil {
		return &pb.ID{Id: id}, nil
	} else {
		return &pb.ID{Id: id}, nil
	}
}

func (s * server) GetEmployee(ctx context.Context, emp *pb.ID) (*pb.Employee, error){
	if empData, err := db.Get(emp.Id); err != nil {
		return nil, err
	} else {
		return empData, nil
	}
}

func (s * server) UpdateEmployee(ctx context.Context, emp *pb.Employee) (*pb.ID, error){
	
	if id, err := db.Update(ctx,emp); err != nil {
		return &pb.ID{Id: id}, nil
	} else {
		return &pb.ID{Id: id}, nil
	}
	
}

func (s * server) DeleteEmployee(ctx context.Context, emp *pb.ID) (*pb.ID, error){

	if id, err := db.Delete(ctx,emp.Id); err != nil {
		return &pb.ID{Id: id}, nil
	} else {
		return &pb.ID{Id: id}, nil
	}
	
}


