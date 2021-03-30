package main

import (
	"os"
	"net"
	"log"
	"context"
	
	pb "github.com/hitesh-sureify/grpc-template/proto"
	"github.com/hitesh-sureify/grpc-template/db"
	"github.com/hitesh-sureify/grpc-template/middleware"
	"github.com/hitesh-sureify/grpc-template/logger"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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

	opts := middleware.GetGrpcMiddlewareOpts()
	s := grpc.NewServer(opts...)

	pb.RegisterEmployeeServiceServer(s, &server{})

	grpc_prometheus.Register(s)
	middleware.RunPrometheusServer()

	if err := logger.Init(-1, "2006-01-02T15:04:05Z07:00"); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}



func (s * server) CreateEmployee(ctx context.Context, emp *pb.Employee) (*pb.ID, error){
	middleware.Incoming_api_req_counter.Add(1)
	if ra, err := db.Insert(ctx,emp); err != nil {
		middleware.Emp_create_fail_counter.Add(1)
		return &pb.ID{}, err
	} else{
		return &pb.ID{Id : int32(ra)}, nil
	}
}

func (s * server) GetEmployee(ctx context.Context, emp *pb.ID) (*pb.Employee, error){
	middleware.Incoming_api_req_counter.Add(1)
	if empData, err := db.Get(emp.Id); err != nil {
		middleware.Emp_get_fail_counter.Add(1)
		return &pb.Employee{}, err
	} else {
		return empData, nil
	}
}

func (s * server) UpdateEmployee(ctx context.Context, emp *pb.Employee) (*pb.ID, error){
	middleware.Incoming_api_req_counter.Add(1)
	if ra, err := db.Update(ctx,emp); err != nil{
		middleware.Emp_update_fail_counter.Add(1)
		return &pb.ID{}, err
	} else{
		return &pb.ID{Id : int32(ra)}, nil
	}
}

func (s * server) DeleteEmployee(ctx context.Context, emp *pb.ID) (*pb.ID, error){
	middleware.Incoming_api_req_counter.Add(1)	
	if ra, err := db.Delete(ctx,emp); err != nil {
		middleware.Emp_delete_fail_counter.Add(1)
		return &pb.ID{}, err
	}else{
		return &pb.ID{Id : int32(ra)}, nil
	}
	
}


