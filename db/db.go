package db

import (
    //"database/sql"
	"strings"
	"fmt"
	"context"
	"os"

	pb "github.com/hitesh-sureify/grpc-template/proto"
	_ "github.com/go-sql-driver/mysql"
	
	pg "github.com/go-pg/pg/v10"
)

type Employee struct{
	Id     int32    `json:"id"`
	Name   string   `json:"name"`
	Dept   string   `json:"dept"`
	Skills string   `json:"skills"`
}

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	options := &pg.Options{
	   User:     os.Getenv("DB_USER"),
	   Password: os.Getenv("DB_PASS"),
	   Addr:     address,
	   Database: os.Getenv("DB_NAME"),
	   PoolSize: 50,
	}
 con = pg.Connect(options)
	if con == nil {
	   fmt.Println("cannot connect to postgres")
	} else{
	fmt.Println("connect to postgres")
	}
 return con
 }

 func SelectDBPost(pg *pg.DB, id pb.ID) {
	emp := &Employee{}
	err := pg.Model(emp).Where("id = ?", id.Id).First()
	fmt.Println(emp)
	fmt.Println(err)
 }

func Get(id int32) (*pb.Employee, error) {

    db := NewDBConn()
	defer db.Close()

	empObj := &Employee{}
	emp := pb.Employee{}

	fmt.Println(id)

    err := db.Model(empObj).Where("id = ?", id).First()
    if err != nil {
        return nil, err
	}

	fmt.Println(empObj)

	emp.Id = empObj.Id
	emp.Name = empObj.Name
	emp.Dept = empObj.Dept
	emp.Skills = strings.Split(empObj.Skills, ",")

	return &emp, nil
}

func Insert(ctx context.Context, emp *pb.Employee) (int, error) {
    db := NewDBConn()
	defer db.Close()

	empObj := &Employee{Id : emp.Id, Dept : emp.Dept, Name : emp.Name, Skills : strings.Join(emp.Skills, ",")}

	res, err := db.Model(empObj).Insert()

	if err != nil {
		return -1, err
	}

	return res.RowsAffected(), nil
}

func Update(ctx context.Context, emp *pb.Employee) (int, error) {


	db := NewDBConn()
	defer db.Close()

	empObj := &Employee{Id : emp.Id, Dept : emp.Dept, Name : emp.Name, Skills : strings.Join(emp.Skills, ",")}

	res, err := db.Model(empObj).Where("id = ?", emp.Id).Update()
    if err != nil {
        return -1, err
	}

	return res.RowsAffected(), nil
}

func Delete(ctx context.Context, id *pb.ID) (int, error) {
    db := NewDBConn()
	defer db.Close()

	empObj := &Employee{}
	
	res, err := db.Model(empObj).Where("id = ?", id.Id).Delete()
	if err != nil {
        return -1, err
	}

	return res.RowsAffected(), nil
}
