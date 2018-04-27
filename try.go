package try4go

type try struct {
	err error
}

//New
func New(fn func()(interface{}, error),out interface{}) try {
	out,err :=fn()
	return try{err}
}
func Empty() try {
	return try{nil}
}

func (t try) Then(fn func()(interface{}, error),out interface{}) try  {
	if t.hasError(){
		return t
	}
	return New(fn,out)
}
func (t try) ThenWithOutCallBack(fn func() error) try  {
	if t.hasError(){
		return t
	}

	return try{fn()}
}
func (t try) OnError(fn func(err error)) {
	if t.hasError(){
		fn(t.err)
	}
}
func (t try)hasError()bool {
	return t.err!=nil
}

func (t try)Retry(attempt int,fn func()(interface{}, error),out interface{}) try {
	if t.hasError(){
		return t
	}
	return t.Then(fn,out).retryHelper(attempt-1,fn,out)
}

func (t try)retryHelper(attempt int,fn func()(interface{}, error),out interface{}) try {
	if !t.hasError()||attempt==0{
		return t
	}
	_,err:=fn()
	t.err=err
	return t.retryHelper(attempt-1,fn,out)

}

