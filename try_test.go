package try4go

import (
	"testing"
	"errors"
	"fmt"
)

func TestNew(t *testing.T) {

	if Empty().err!=nil{
		t.Fatal("a new try's err should be nil.")
	}

}
func TestHasError(t *testing.T) {

	try:=Empty()

	if try.hasError(){
		t.Fatal("a new try's err shouldn't has error.")
	}
	try.err=errors.New("test err")
	if !try.hasError(){
		t.Fatal("the try shouldn't has error.")
	}

}
func TestThen(t *testing.T) {


	var isSuccess bool
	try:=Empty()
	newTry :=try.Then(func(interface{}) (interface{}, error) {
		return false,errors.New("err1")
	}).Then(func(interface{}) (interface{}, error) {
		return true,errors.New("err1")
	})

	if !newTry.hasError(){
		t.Fatal("the try should has error.")
	}
	if isSuccess{
		t.Fatal("isSuccess should be false")
	}

}
func TestRetry3Times(t *testing.T) {


	var attempt int
	try:=Empty()

	try.Retry(3,func(interface{}) (interface{}, error) {
		attempt++

		return attempt,errors.New("err1")
	})

	fmt.Println(attempt)
	if attempt!=3{
		t.Fatal("attempt should be 3")
	}


}
func TestRetry3TimesButLastTimeSuccess(t *testing.T) {


	var attempt int
	try:=Empty()

	r:=try.Retry(3,func(interface{}) (interface{}, error) {
		var err error
		if attempt!=3{
			err=errors.New("err")
		}

		return attempt,err
	})

	fmt.Println(attempt)
	if !r.hasError(){
		t.Fatal("the try should have error")
	}

}
func TestRetry3TimesButAlreadyGotErrorBeforeRetryOpt(t *testing.T) {


	var attempt int
	try:=Empty()


	try.err=errors.New("err")
	try.Retry(3,func(interface{}) (interface{}, error) {


		return attempt,nil
	})

	fmt.Println(attempt)
	if !try.hasError(){
		t.Fatal("the try should have erro")
	}

}

func TestOnError(t *testing.T) {
	New(func() (interface{}, error) {
		return nil,errors.New("err")
	}).OnError(func(err error) {
		fmt.Println(err)
	})
}

//func TestPerformanceBenchmark(t *testing.T)  {
//	var times int
//	var start time.Time
//	var elapse int
//	f1:= func() {
//		i,err:=strconv.Atoi("1")
//		if err==nil{
//			times+=i
//		}
//	}
//	f2:= func() {
//		New(func() (interface{}, error) {
//			return strconv.Atoi("1")
//		}).OnSuccess(func(i interface{}) {
//			times+=i.(int)
//		})
//	}
//
//
//	start=time.Now()
//	for i:=0;i<100;i++{
//		f1()
//	}
//	log.Println("======go=====")
//	elapse=time.Now().Nanosecond()-start.Nanosecond()
//	log.Printf("%v ns",elapse)
//
//	times=0
//	start=time.Now()
//	for i:=0;i<100;i++{
//		f2()
//	}
//	log.Println("======try4go=====")
//	elapse=time.Now().Nanosecond()-start.Nanosecond()
//	log.Printf("%v ns",elapse)
//
//
//}

