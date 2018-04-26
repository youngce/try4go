package try4go

import (
	"testing"
	"errors"
	"fmt"
)

func TestNew(t *testing.T) {

	if New().err!=nil{
		t.Fatal("a new try's err should be nil.")
	}

}
func TestHasError(t *testing.T) {

	try:=New()

	if try.hasError(){
		t.Fatal("a new try's err shouldn't has error.")
	}
	try.err=errors.New("test err")
	if try.hasError(){
		t.Fatal("the try should has error.")
	}

}
func TestThen(t *testing.T) {


	var isSuccess bool
	try:=New()
	try.Then(func() (interface{}, error) {
		return false,errors.New("err1")
	},isSuccess).Then(func() (interface{}, error) {
		return true,errors.New("err1")
	},isSuccess)

	if !try.hasError(){
		t.Fatal("the try should has error.")
	}
	if isSuccess{
		t.Fatal("isSuccess should be false")
	}

}
func TestRetry3Times(t *testing.T) {


	var attempt int
	try:=New()

	try.Retry(3,func() (interface{}, error) {
		attempt++

		return attempt,errors.New("err1")
	},attempt)

	fmt.Println(attempt)
	if attempt!=3{
		t.Fatal("attempt should be 3")
	}


}
func TestRetry3TimesButLastTimeSuccess(t *testing.T) {


	var attempt int
	try:=New()

	try.Retry(3,func() (interface{}, error) {
		var err error
		if attempt!=3{
			err=errors.New("err")
		}

		return attempt,err
	},attempt)

	fmt.Println(attempt)
	if !try.hasError(){
		t.Fatal("the try should have error")
	}

}
func TestRetry3TimesButAlreadyGotErrorBeforeRetryOpt(t *testing.T) {


	var attempt int
	try:=New()


	try.err=errors.New("err")
	try.Retry(3,func() (interface{}, error) {


		return attempt,nil
	},attempt)

	fmt.Println(attempt)
	if !try.hasError(){
		t.Fatal("the try should have erro")
	}

}
