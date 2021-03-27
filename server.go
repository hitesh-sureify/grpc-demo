package main

import (
	"fmt"
	"net"
	"log"
	"context"
	
	pb "github.com/hitesh-sureify/GrpcDemo/proto"
	"github.com/hitesh-sureify/GrpcDemo/db"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedEmployeeServiceServer
}

func main() {
	
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
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
	db.Insert(emp)
	fmt.Println("emp created")
	return &pb.ID{Id: 0}, nil
}

func (s * server) GetEmployee(ctx context.Context, emp *pb.ID) (*pb.Employee, error){
	Get(emp.Id)
	fmt.Println("got employee")
	return nil, nil
}

func (s * server) UpdateEmployee(ctx context.Context, emp *pb.Employee) (*pb.ID, error){
	Update(emp)
	fmt.Println("update employee")
	return &pb.ID{Id: 0}, nil
}

func (s * server) DeleteEmployee(ctx context.Context, emp *pb.ID) (*pb.ID, error){
	Delete(emp.Id)
	fmt.Println("emp deleted")
	return &pb.ID{Id: 0}, nil
}

// func (s * server) mustEmbedUnimplementedEmployeeServiceServer(){
// 	fmt.Println("just an implementation")
// }

