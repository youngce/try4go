package try4go

type try struct {
	err error
}

//New
func New() *try {
	return &try{nil}
}
func (t *try) Then(fn func()(interface{}, error),out interface{}) *try  {
	if t.hasError(){
		return t
	}
	out,err :=fn()
	t.err=err
	return t
}
func (t *try) OnError(fn func(err error)) {
	fn(t.err)

}
func (t *try)hasError()bool {
	return t.err!=nil
}

func (t *try)Retry(attempt int,fn func()(interface{}, error),out interface{}) *try {
	if t.hasError(){
		return t
	}
	return t.Then(fn,out).retryHelper(3,fn,out)
}

func (t *try)retryHelper(attempt int,fn func()(interface{}, error),out interface{}) *try {
	if !t.hasError(){
		return t
	}
	_,err:=fn()
	t.err=err
	return t.retryHelper(attempt-1,fn,out)

}

