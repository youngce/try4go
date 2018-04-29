package main

import (
	"try4go"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/gommon/log"
)

func main() {

	txTried := try4go.New(func() (interface{}, error) {
		return sql.Open("sqlite3", "./foo.db")
	}).Then(func(in interface{}) (interface{}, error) {
		return in.(*sql.DB).Begin()
	})


	execTried:=txTried.Then(func(in interface{}) (interface{}, error) {
		return in.(*sql.Tx).Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
	}).Then(func(in interface{}) (interface{}, error) {
		id:=4
		return in.(*sql.Stmt).Exec(id)
	})
	//execTried.Merge(txTried, func(kv try4go.Tuple2) (interface{}, error) {
	//	tx:=kv.V2.(*sql.Tx)
	//	return tx,tx.Commit()
	//})
	
	execTried.OnError(func(err error) {
		log.Error(err)
		txTried.OnSuccess(func(i interface{}) {
			i.(*sql.Tx).Rollback()
		})
	})
	execTried.OnSuccess(func(i interface{}) {
		txTried.OnSuccess(func(i interface{}) {
			i.(*sql.Tx).Commit()
		})
	})




}
