package main

import(
	"log"
	"context"
	"net/http"
	"os"
	"strings"
	"time"
	"fmt"
	"strconv"
	"encoding/json"

	pb "github.com/hitesh-sureify/grpc-demo/proto"
	"github.com/hitesh-sureify/grpc-demo/middleware"

	"google.golang.org/grpc"
	"github.com/gorilla/mux"
)

var c pb.EmployeeServiceClient

type EmployeeAPI struct{
	Id     int32    `json:"id"`
	Name   string   `json:"name"`
	Dept   string   `json:"dept"`
	Skills string   `json:"skills"`
}


func main(){
	conn, err := grpc.Dial(os.Getenv("GRPC_SRV_ADDR"), grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to the server")
	}

	defer conn.Close()

	c = pb.NewEmployeeServiceClient(conn)

	r := mux.NewRouter()

	r.HandleFunc("/api/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/employees", createEmployee).Methods("POST")
	r.HandleFunc("/api/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", deleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
	
}

func getEmployee(w http.ResponseWriter, r *http.Request) {


	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	empId, _ := strconv.Atoi(params["id"])

	var msg string

	empData, err := c.GetEmployee(ctx, &pb.ID{Id : int32(empId)})
	if err != nil{
		msg = fmt.Sprintf("Could not get employee : %s", err.Error())
	} else{
		msg = fmt.Sprintf("Employee record fetched for Id %d =>  Name : %s, Dept : %s, Skills : %s", empId, empData.Name, empData.Dept, strings.Join(empData.Skills, ","))
	}
	json.NewEncoder(w).Encode(msg)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {


	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	var emp EmployeeAPI
	_ = json.NewDecoder(r.Body).Decode(&emp)

	var msg string

	empData, err := c.CreateEmployee(ctx, &pb.Employee{Name: emp.Name, Dept: emp.Dept, Skills: strings.Split(emp.Skills, ",")})
	if err != nil{
		msg = fmt.Sprintf("Could not create employee record : %s", err.Error())
	} else{
		msg = fmt.Sprintf("Employee created  with ID : %d", empData.Id)
	}

	json.NewEncoder(w).Encode(msg)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var emp EmployeeAPI
	_ = json.NewDecoder(r.Body).Decode(&emp)

	var msg string
	empId, _ := strconv.Atoi(params["id"])

	empData, err := c.UpdateEmployee(ctx, &pb.Employee{Id: int32(empId), Name: emp.Name, Dept: emp.Dept, Skills: strings.Split(emp.Skills, ",")})
	if err != nil{
		msg = fmt.Sprintf("Could not update employee record : %s", err.Error())
	}
	if empData.Id < 0{
		msg = fmt.Sprintf("could not update employee record")
	} else{
		msg = fmt.Sprintf("Employee record updated. Rows affected : %d", empData.Id)
	}

	json.NewEncoder(w).Encode(msg)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	empId, _ := strconv.Atoi(params["id"])

	var msg string

	empData, err := c.DeleteEmployee(ctx, &pb.ID{Id : int32(empId)})
	if err != nil{
		msg = fmt.Sprintf("Could not delete employee record : %s", err.Error())
	}
	if empData.Id <= 0{
		msg = fmt.Sprintf("Could not delete employee record")
	} else {
		msg = fmt.Sprintf("Employee record deleted. Rows affected : %d", empData.Id)
	}

	json.NewEncoder(w).Encode(msg)
}