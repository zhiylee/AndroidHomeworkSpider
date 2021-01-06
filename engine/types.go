package engine

import (
	"io"
)

type Request struct {
	Url string
	ParserFunc func(r io.Reader,request Request) (ParserRes,error)
	Data interface{}
}

type ParserRes struct {
	Requests []Request
	Items []interface{}
}

type WorkerQ []Worker

func (q *WorkerQ) push(v Worker)  {
	*q = append(*q,v)
}

func (q *WorkerQ) pop() Worker {
	v:=(*q)[0]
	*q = (*q)[1:]
	return v
}

type RequestQ []Request

func (q *RequestQ) push(v ...Request)  {
	*q = append(*q,v...)
}

func (q *RequestQ) pop() Request {
	v:=(*q)[0]
	*q = (*q)[1:]
	return v
}

func NilParser(r io.Reader,data interface{}) (ParserRes,error){
	return ParserRes{},nil
}