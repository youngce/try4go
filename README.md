# 簡介
如果你覺得go原生的錯誤處理方式不夠優美，或許可以考慮使用try4go.

try4go設計想法來自於scala中的[Try](https://www.scala-lang.org/api/2.12.4/scala/util/Try.html), 

但遺憾的是...go中並沒有```generic```，無法實現出Scala中那樣優美的操作，但try4go仍然改善error handling處理方式。

（以下為0.0.x版本）

你可以參考以下的範例代碼:

golang
```go
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        panic(err)
    }
    // create table
    _,err:=db.Exec("CREATE TABLE `userinfo` (" +
        			"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
        			"`username` VARCHAR(64) NULL," +
        			"`departname` VARCHAR(64) NULL," +
        			"`created` DATE NULL);")
    if err != nil {
        panic(err)
    }

    // insert
    stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
    if err != nil {
        panic(err)
    }

    res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
    if err != nil {
        panic(err)
    }

    id, err := res.LastInsertId()
    if err != nil {
        panic(err)
    }

    fmt.Println(id)
```

try4go
```go
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

```

當你查看try4go中的Then方法可以發現當```try```中有error時，將不會繼續執行後續的函數。
# Benchmark
在```benchmark_test.go```中你可以找到測試的程式碼，try4go效能大概慢了原生go錯誤處理3.5倍，

主要原因在於try4go每次都需要做一次強轉型(Parse)，若不使用強轉型
```
BenchmarkTry4go-8                       20000000                67.0 ns/op
BenchmarkPureGo-8                       100000000               17.6 ns/op
BenchmarkTry4goWithoutParse-8           10000000                121 ns/op

```
# Change Log

## [0.1.0] - 2018-04-27

### Changed
- API改為Lambda風格
- OnSuccess() API


### Removed

## [0.0.2] - 2018-04-27
### Added
- 增加```ThenWithOutCallBack```
- ```Empty（）```取代原本的```New()```

### Changed
- 修改```New```參數，與```Then```相同
### Fixed
- ```OnError```無```err==nil```不執行

### Removed