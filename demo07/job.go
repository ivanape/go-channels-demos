package main

import (
	"errors"
	"reflect"
	"runtime"
)

var (
	ErrTimeFormat           = errors.New("time format error")
	ErrParamsNotAdapted     = errors.New("the number of params is not adapted")
	ErrNotAFunction         = errors.New("only functions can be schedule into the job queue")
	ErrPeriodNotSpecified   = errors.New("unspecified job period")
	ErrParameterCannotBeNil = errors.New("nil paramaters cannot be used with reflection")
)

type Status int

type Job struct {
	id      string
	status  Status
	result  JobResult
	funcs   map[string]interface{}   // Map for the function task store
	fparams map[string][]interface{} // Map for function and  params of function
	jobFunc string
}

type JobResult struct {
	message string
	err     error
}

const (
	Pending Status = iota
	Running
	Done
	Cancelled
)

func (j *Job) SetTask(jobFun interface{}, params ...interface{}) error {
	typ := reflect.TypeOf(jobFun)
	if typ.Kind() != reflect.Func {
		return ErrNotAFunction
	}
	fname := getFunctionName(jobFun)
	j.funcs[fname] = jobFun
	j.fparams[fname] = params
	j.jobFunc = fname

	return nil
}

func (j *Job) Run() ([]reflect.Value, error) {
	result, err := callJobFuncWithParams(j.funcs[j.jobFunc], j.fparams[j.jobFunc])
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func callJobFuncWithParams(jobFunc interface{}, params []interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(jobFunc)
	if len(params) != f.Type().NumIn() {
		return nil, ErrParamsNotAdapted
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in), nil
}
