package try4go

type Tuple2 struct {
	V1 interface{}
	V2 interface{}
}




type try struct {
	err error
	succ interface{}
}
type tryOp func(in interface{}) (interface{},error)

//New
func New(fn func()(interface{}, error)) try {

	v,e:=fn()

	return try{e,v}
}
func Empty() try {
	return try{nil,nil}
}

func (t1 try) Then(op tryOp) try  {
	if t1.hasError(){
		return t1
	}


	return New(func() (interface{}, error) {
		return op(t1.succ)
	})
}
func (t1 try) Err() error  {
	return t1.err
}
func (t1 try) RecoverError(fn func(error) error) try  {
	if t1.hasError(){
		return New(func() (interface{}, error) {
			return t1.succ, fn(t1.err)
		})
	}
	return t1
	//return t1.err
}
func (t1 try) Merge(t2 try,op func(kv Tuple2) (interface{},error)) try{
	if t1.hasError(){
		return t1
	}
	if t2.hasError(){
		return t2
	}
	return New(func() (interface{}, error) {
		return op(Tuple2{t1.succ,t2.succ})
	})
}

func (t1 try) OnError(fn func(err error)) {
	if t1.hasError(){
		fn(t1.err)
	}
}
func (t1 try) Result() (interface{},error) {
	return t1.succ,t1.err
}

func (t1 try)OnSuccess(fn func(interface{}))  {
	if !t1.hasError(){
		fn(t1.succ)
	}
}


func (t1 try)hasError()bool {
	return t1.err!=nil
}
//
func (t1 try)Retry(attempt int,op tryOp) try {
	if t1.hasError(){
		return t1
	}
	return t1.Then(op).retryHelper(attempt-1,op)
}
//
func (t1 try)retryHelper(attempt int,op tryOp) try {
	if !t1.hasError()||attempt==0{
		return t1
	}
	_,err:=op(t1.succ)
	t1.err=err
	return t1.retryHelper(attempt-1,op)

}

