package main

import (
	"database/sql"
	_ "github.com/mikespook/mymysql/godrv"
	"log"
)

func main() {
	// 使用 mymysql 驱动打开一个 sql.DB。
	// 对于 mymysql 中，dsn 有多种写法。
	// 使用 tcp 协议：[tcp://addr/]dbname/user/password[?params]
	// 使用 unix sock：[unix://sockpath/]dbname/user/password[?params]
	// 括号中是可选的内容，当协议信息未指定时，默认使用 tcp://127.0.0.1:3306/。
	// params 部分可设置两个参数
	//   charset – 用于 ‘set names’ 设置连接编码。
	//   keepalive – 每 keepalive 秒向服务器发送 PING。
	//
	// 如果密码含有斜线（/），则用星号（*）代替。
	// 如果密码含有星号（*），用两个星号（**）代替。
	//   pass/wd => pass*wd
	//   pass*wd => pass**wd
	db, err := sql.Open("mymysql", "tcp://127.0.0.1:3306/paths/root/nice?charset=utf8&keepalive=1200")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// select
	rows, err := db.Query("select id, code from `org_job`")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var code, id string
		rows.Scan(&code, &id)
		log.Printf("[ROWS]key=%s, value=%s", code, id)
	}

	// select where
	row := db.QueryRow("select * from `test` where `key` = ?", "name")
	var k, v string
	row.Scan(&k, &v)
	log.Printf("[ROW]key=%s, value=%s", k, v)

	// delete
	rslt, err = db.Exec("delete from `test`")
	if a, err := rslt.RowsAffected(); err != nil {
		log.Print(err)
	} else {
		log.Printf("[DEL]Affected rows=%d", a)
	}
	if id, err := rslt.LastInsertId(); err != nil {
		log.Print(err)
	} else {
		log.Printf("[DEL]Last insert id=%d", id)
	}

	// insert
	stmt, err := db.Prepare("insert into `test` (`key`, `value`) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rslt, err := stmt.Exec("name", "foobar")
	if err != nil {
		log.Fatal(err)
	}
	if a, err := rslt.RowsAffected(); err != nil {
		log.Print(err)
	} else {
		log.Printf("[INS]Affected rows=%d", a)
	}
	if id, err := rslt.LastInsertId(); err != nil {
		log.Print(err)
	} else {
		log.Printf("[INS]Last insert id=%d", id)
	}
}
