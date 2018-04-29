package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"try4go"
	"log"
)

func main() {

	try:=try4go.New(func() (interface{}, error) {
		return sql.Open("sqlite3", "./foo.db")
	}).Then(func(in interface{}) (interface{}, error) {
		// Create userinfo table
		_,err:=in.(*sql.DB).Exec("CREATE TABLE `userinfo` (" +
			"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
			"`username` VARCHAR(64) NULL," +
			"`departname` VARCHAR(64) NULL," +
			"`created` DATE NULL);")
		return in,err
	}).Then(func(in interface{}) (interface{}, error) {
		// Insert
		return in.(*sql.DB).Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	}).Then(func(in interface{}) (interface{}, error) {
		return in.(*sql.Stmt).Exec("astaxie", "研发部门", "2012-12-09")
	}).Then(func(in interface{}) (interface{}, error) {
		// Get Last InsertId
		return in.(sql.Result).LastInsertId()
	})
	try.OnError(func(err error) {
		log.Fatal(err)
	})
	try.OnSuccess(func(i interface{}) {
		log.Println(i)
	})
	//checkErr(err)
	//
	//// insert
	//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	//checkErr(err)
	//
	//res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	//checkErr(err)
	//
	//id, err := res.LastInsertId()
	//checkErr(err)
	//
	//fmt.Println(id)
	//// update
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("astaxieupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)

	//fmt.Println(affect)
}
