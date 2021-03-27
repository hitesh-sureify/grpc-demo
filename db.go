package main

import (
    "database/sql"
	"strings"
	"fmt"
	"context"

	pb "github.com/hitesh-sureify/GrpcDemo/proto"
    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "hitesh"
    dbPass := "68c#sistEdgCD4"
	dbName := "MYSQLTEST"
	dbHost := "172.17.0.1:3306"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(" + dbHost + ")/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func Get(id int32) (*pb.Employee, error) {

    db := dbConn()
	defer db.Close()

	var emp pb.Employee
	var skills string

    err := db.QueryRow("SELECT name, dept, skills FROM employees WHERE id=?;", id).Scan(&emp.Name, &emp.Dept, &skills)
    if err != nil {
        return nil, err
	}
	emp.Skills = strings.Split(skills, ",")

	return &emp, nil
}

func Insert(ctx context.Context, emp *pb.Employee) (int32,error) {
    db := dbConn()
	defer db.Close()

	stmt := fmt.Sprintf("insert into employees(name, dept, skills) values ('%s', '%s', '%s');", emp.Name, emp.Dept, strings.Join(emp.Skills, ","))
	
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}


	res, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int32(lastInsertId), nil
}

func Update(ctx context.Context, emp *pb.Employee) (int32,error) {


	db := dbConn()
	defer db.Close()

	stmt := fmt.Sprintf("update employees set name='%s', dept='%s', skills='%s' WHERE id='%d';", emp.Name, emp.Dept, strings.Join(emp.Skills, ","), emp.Id)
	
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}


	res, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int32(count), nil
}

func Delete(ctx context.Context, id int32) (int32,error) {
    db := dbConn()
	defer db.Close()
	
	stmt := fmt.Sprintf("delete from employees where id = '%d';", id)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

    res, err := db.ExecContext(ctx, stmt)
    if err != nil {
		fmt.Println(err)
        tx.Rollback()
		return -1, err
	}
	
	count, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	
    err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int32(count), nil
}
