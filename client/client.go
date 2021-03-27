package main

import(
	"log"
	"context"
	"bufio"
	"os"
	"strings"
	"time"
	"fmt"
	"strconv"

	pb "github.com/hitesh-sureify/grpc-demo/proto"

	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to the server")
	}

	defer conn.Close()

	c := pb.NewEmployeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	fmt.Println("Grpc Employee Service app!!!!!!!!")

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	switch(text){
	case "1\n":
		fmt.Println("\nEnter Name : ")
		name := bufio.NewReader(os.Stdin)
		n, _ := name.ReadString('\n')
		n = strings.Trim(n, "\n")

		fmt.Println("\nEnter Dept : ")
		dept := bufio.NewReader(os.Stdin)
		d, _ := dept.ReadString('\n')
		d = strings.Trim(d, "\n")

		fmt.Println("\nEnter Skills comma separated : ")
		skills := bufio.NewReader(os.Stdin)
		s, _ := skills.ReadString('\n')
		s = strings.Trim(s, "\n")

		emp, err := c.CreateEmployee(ctx, &pb.Employee{Name: n, Dept: d, Skills: strings.Split(s, ",")})
		if err != nil{
			log.Fatalf("could not add employee : %v", err)
		}
		fmt.Println("Employee added \n with ID : ", emp.Id)
	
	case "2\n":
		fmt.Println("\nEnter ID : ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		empId, _ := strconv.Atoi(strings.Trim(i, "\n"))

		emp, err := c.GetEmployee(ctx, &pb.ID{Id : int32(empId)})
		if err != nil{
			log.Fatalf("could not get employee : %v", err)
		}
		fmt.Println("Employee Name : ", emp.Name)
		fmt.Println("Employee Dept : ", emp.Dept)
		fmt.Println("Employee Skills : ", strings.Join(emp.Skills, ","))
	
	case "3\n":
		fmt.Println("\nEnter ID : ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		empId, _ := strconv.Atoi(strings.Trim(i, "\n"))

		fmt.Println("\nEnter Name : ")
		name := bufio.NewReader(os.Stdin)
		n, _ := name.ReadString('\n')
		n = strings.Trim(n, "\n")

		fmt.Println("\nEnter Dept : ")
		dept := bufio.NewReader(os.Stdin)
		d, _ := dept.ReadString('\n')
		d = strings.Trim(d, "\n")

		fmt.Println("\nEnter Skills comma separated : ")
		skills := bufio.NewReader(os.Stdin)
		s, _ := skills.ReadString('\n')
		s = strings.Trim(s, "\n")

		emp, err := c.UpdateEmployee(ctx, &pb.Employee{Id: int32(empId), Name: n, Dept: d, Skills: strings.Split(s, ",")})
		if err != nil{
			log.Fatalf("could not update employee : %v", err)
		}
		if emp.Id < 0{
			log.Fatalf("could not update employee record")
		}
		fmt.Println("Employee updated. Rows affected : ", emp.Id)
	
	case "4\n":
		fmt.Println("\nEnter ID : ")
		id := bufio.NewReader(os.Stdin)
		i, _ := id.ReadString('\n')
		empId, _ := strconv.Atoi(strings.Trim(i, "\n"))

		emp, err := c.DeleteEmployee(ctx, &pb.ID{Id : int32(empId)})
		if err != nil{
			log.Fatalf("could not delete employee record : %v", err)
		}
		if emp.Id < 0{
			log.Fatalf("could not delete employee record")
		}
		fmt.Println("Employee record deleted. Rows affected : ", emp.Id)

	}
}