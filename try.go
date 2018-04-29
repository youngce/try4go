package try4go

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

func (t try) Then(op tryOp) try  {
	if t.hasError(){
		return t
	}

	return New(func() (interface{}, error) {
		return op(t.succ)
	})
}

func (t try) OnError(fn func(err error)) {
	if t.hasError(){
		fn(t.err)
	}
}
func (t try)OnSuccess(fn func(interface{}))  {
	if !t.hasError(){
		fn(t.succ)
	}
}
func (t try)hasError()bool {
	return t.err!=nil
}
//
func (t try)Retry(attempt int,op tryOp) try {
	if t.hasError(){
		return t
	}
	return t.Then(op).retryHelper(attempt-1,op)
}
//
func (t try)retryHelper(attempt int,op tryOp) try {
	if !t.hasError()||attempt==0{
		return t
	}
	_,err:=op(t.succ)
	t.err=err
	return t.retryHelper(attempt-1,op)

}

