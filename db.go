package db

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

	pb "github.com/hitesh-sureify/GrpcDemo/proto"
    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "hitesh"
    dbPass := "68c#sistEdgCD4"
	dbName := "MYSQLTEST"
	dbHost := "172.17.0.1:3306"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(" + dbHost ")/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func Get(id int32) {
    db := dbConn()
    defer db.Close()
    empDB, err := db.Query("SELECT * FROM Employee WHERE id=?", id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(empDB)
}

func Insert(emp *pb.Employee) {
    db := dbConn()
    defer db.Close()
    insForm, err := db.Prepare("UPDATE Employee SET name=?, dept=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	id, _ := insForm.Exec(emp.Name, emp.Dept)
	fmt.Println(id)
}

func Update(emp *pb.Employee) {
    db := dbConn()
    defer db.Close()
    insForm, err := db.Prepare("UPDATE Employee SET name=?, dept=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	id, _ := insForm.Exec(emp.Name, emp.Dept, Emp.Id)
	fmt.Println(id)
}

func Delete(id int32) {
    db := dbConn()
    defer db.Close()
    delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    data, _ := delForm.Exec(id)
    fmt.Println(data)
}
