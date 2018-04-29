package try4go

import (
	"testing"
	"strconv"
)

func BenchmarkTry4goWithoutParse(b *testing.B) {
	var times int

	//var try try
	try:=New(func() (interface{}, error) {
		return strconv.Atoi("1")	})
	for n := 0; n < b.N; n++ {
		try.Then(func(_ interface{}) (interface{}, error) {
			times+=1
			return strconv.Atoi(strconv.Itoa(times))
		})
	}
	//fmt.Println(times)
}
func BenchmarkTry4go(b *testing.B) {
	var times int
	//var try try
	try:=New(func() (interface{}, error) {
		return strconv.Atoi("1")
	})
	for n := 0; n < b.N; n++ {
		try.Then(func(in interface{}) (interface{}, error) {
			times+=in.(int)
			return strconv.Atoi(strconv.Itoa(times))
		})
	}
	//fmt.Println(times)
}
func BenchmarkPureGo(b *testing.B) {
	var times int
	var i int
	var err error
	i,err=strconv.Atoi("1")
	for n := 0; n < b.N; n++ {
		if err==nil{
			times+=i
			i,err=strconv.Atoi(strconv.Itoa(times))
		}
	}

}


