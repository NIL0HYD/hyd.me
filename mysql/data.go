package main

import (
	"database/sql"
	_ "github.com/mikespook/mymysql/godrv"
	"log"
)

type Job struct {
	Id     int
	Code   string
	Name   string
	Office int
	Parent int
	Level  int
	Duty   string
	Order  int
}

func getDB() (db *sql.DB) {
	db, err := sql.Open("mymysql", "tcp://127.0.0.1:3306/paths/root/nice?charset=utf8&keepalive=1200")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	//Select()
	//Update()
	Insert()
	Select()
}

func Select() {
	db := getDB()
	defer db.Close()
	// select
	rows, err := db.Query("select * from `org_job`")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		job := Job{}
		err := rows.Scan(&job.Id, &job.Code, &job.Name,
			&job.Office, &job.Parent, &job.Level,
			&job.Duty, &job.Order)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("[ROWS]key=%v, value=%v, level=%v ", job.Id, job.Code, job.Level)
	}
}

func Update() {
	db := getDB()
	defer db.Close()

	stmt, err := db.Prepare("update org_job set level=? where id=?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	rslt, err := stmt.Exec(2, 1)
	if a, err := rslt.RowsAffected(); err != nil {
		log.Print(err)
	} else {
		log.Printf("[INS]Affected rows=%d", a)
	}
}

func Insert() {
	db := getDB()
	defer db.Close()

	job := Job{}
	job.Code = "101010"
	job.Name = "Test"
	job.Level = 1
	job.Office = 1
	job.Order = 3

	stmt, err := db.Prepare("insert into org_job(code, name, office, parent, level, duty, `order`) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()
	rslt, err := stmt.Exec(job.Code, job.Name, job.Office, job.Parent, job.Level, job.Duty, job.Order)
	if _, err := rslt.RowsAffected(); err != nil {
		log.Fatal(err)
		return
	}
}
