# Change Log

## [0.0.2] - 2018-04-27
### Added
- 增加```ThenWithOutCallBack```
- ```Empty（）```取代原本的```New()```

### Changed
- 修改```New```參數，與```Then```相同
### Fixed
- ```OnError```無```err==nil```不執行

### Removed


# 簡介
如果你覺得go原生的錯誤處理方式不夠優美，或許可以考慮使用try4go.

try4go設計想法來自於scala中的[Try](https://www.scala-lang.org/api/2.12.4/scala/util/Try.html), 

但遺憾的是...go中並沒有```generic```，無法實現出Scala中那樣優美的操作，但try4go仍然改善error handling處理方式。

你可以參考以下的範例代碼:

golang
```go
    db,err:=sql.Open(...)
    if err!=nil{
        rows,queryErr:=db.Query("your query statement1")
        if queryErr!=nil{
            scanErr:=rows.Scan(...)
            if scanErr!=nil{
                ....
            }
        }
    }
```

try4go
```go
    var db *sql.DB
    var rows *sql.Rows
    try4go.New().Then(func() (interface{}, error) {
		return sql.Open(...)
	},db).Then(func() (interface{}, error) {
		return db.Query("your query statement1")
	},rows).Then(func() (interface{}, error) {
		return nil,err
	},nil)
```

當你查看try4go中的Then方法可以發現當```try```中有error時，將不會繼續執行後續的函數。

另外, try4go也提供```OnError```的Api，方便處理錯誤訊息，如寫Log
```go
try4go.New()
    .Then(...)
    .Then(...)
    .OnError(func(err error) {
    		log.Fatal(err.Error())
    	})
```
