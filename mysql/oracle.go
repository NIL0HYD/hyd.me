package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"os"
)

func oraDB() (db *sql.DB) {
	db, err := sql.Open("oci8", "pi/pi@rd_10.10.71.125")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func main() {
	// 
	os.Setenv("NLS_LANG", "")
	//fmt.Println(os.Getenv("NLS_LANG"))
	//oraSelect()
	//oraInsert()
	oraUpdate()
}

func oraSelect() {
	db := oraDB()
	defer db.Close()
	rows, err := db.Query("select pkid_, name_, key_ from BPM_PROC_DESC")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var pkid string
		var name string
		var key string
		err := rows.Scan(&pkid, &name, &key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(pkid, name, key)
	}
}

func oraUpdate() {
	db := oraDB()
	defer db.Close()
	stmt, err := db.Prepare("update CM_SERIAL_NUMBER set prefix_=:prefix where id_=:id")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	rslt, err := stmt.Exec("CYMCO", 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := rslt.RowsAffected(); err != nil {
		fmt.Println(err)
		return
	}
}

func oraInsert() {
	db := oraDB()
	defer db.Close()
	//You can't use ? to bind parameters. Use :foo like variable.
	stmt, err := db.Prepare("insert into CM_SERIAL_NUMBER(id_, year_) values(:id, :year)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rslt, err := stmt.Exec(7, "2010")
	if err != nil {
		fmt.Println(err)
		return
	}
	if a, err := rslt.RowsAffected(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(a)
	}
}

func oraInsert2() {
	db := oraDB()
	defer db.Close()
	//You can't use ? to bind parameters. Use :foo like variable.
	stmt, err := db.Prepare("insert into CM_SERIAL_NUMBER values(:id, :prefix, :year, :count)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rslt, err := stmt.Exec(2, "CYMCO", "2013", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := rslt.LastInsertId(); err != nil {
		fmt.Println(err)
		return
	}
}
