package db

import (
	"strings"
	"fmt"
	"context"
	"os"

	pb "github.com/hitesh-sureify/grpc-template/proto"
	"github.com/hitesh-sureify/grpc-template/logger"
	
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
		logger.Log.Warn("db connection failed...")
	} else{
		logger.Log.Info("connected to database...")
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
	

    err := db.Model(empObj).Where("id = ?", id).First()
    if err != nil {
		logger.Log.Error("failed to fetch employee data from db, reason : " + err.Error())
        return nil, err
	}
	

	emp.Id = empObj.Id
	emp.Name = empObj.Name
	emp.Dept = empObj.Dept
	emp.Skills = strings.Split(empObj.Skills, ",")

	logger.Log.Info("employee record fetched from db.")

	return &emp, nil
}

func Insert(ctx context.Context, emp *pb.Employee) (int, error) {
    db := NewDBConn()
	defer db.Close()

	empObj := &Employee{Id : emp.Id, Dept : emp.Dept, Name : emp.Name, Skills : strings.Join(emp.Skills, ",")}

	res, err := db.Model(empObj).Insert()

	if err != nil {
		logger.Log.Error("failed to insert employee data into db, reason : " + err.Error())
		return -1, err
	}

	logger.Log.Info("employee record inserted into db.")

	return res.RowsAffected(), nil
}

func Update(ctx context.Context, emp *pb.Employee) (int, error) {


	db := NewDBConn()
	defer db.Close()

	empObj := &Employee{Id : emp.Id, Dept : emp.Dept, Name : emp.Name, Skills : strings.Join(emp.Skills, ",")}

	res, err := db.Model(empObj).Where("id = ?", emp.Id).Update()
    if err != nil {
		logger.Log.Error("failed to update employee data in db, reason : " + err.Error())
        return -1, err
	}

	logger.Log.Info("employee record updated in db.")

	return res.RowsAffected(), nil
}

func Delete(ctx context.Context, id *pb.ID) (int, error) {
    db := NewDBConn()
	defer db.Close()

	empObj := &Employee{}
	
	res, err := db.Model(empObj).Where("id = ?", id.Id).Delete()
	if err != nil {
		logger.Log.Error("failed to delete employee data from db, reason : " + err.Error())
        return -1, err
	}

	logger.Log.Info("employee record deleted from db.")

	return res.RowsAffected(), nil
}
