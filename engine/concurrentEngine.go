package engine

import (
	"androidHomeworkSpider/model"
	"fmt"
	"log"
	"os"
	"time"
)

type Concurrent struct{
	WorkerCount int

	RequestQ RequestQ  // tasks waiting handle
	WorkerQ WorkerQ // available workers queue

	RequestChan chan []Request
	WorkerChan chan Worker
	ParserResChan chan ParserRes
}

func (e *Concurrent) Run(seeds ...Request)  {
	// init queue
	e.RequestQ = []Request{}
	e.WorkerQ = WorkerQ{}

	e.RequestChan = make(chan []Request)
	e.ParserResChan = make(chan ParserRes,e.WorkerCount)
	e.WorkerChan = make(chan Worker,e.WorkerCount)

	// creat workers
	for i:=0;i<e.WorkerCount;i++{
		e.WorkerQ.push(e.CreateWorker(i,&e.ParserResChan))
	}

	for _,seed:= range seeds {
		e.RequestQ.push(seed)
	}

	// scheduler
	go func() {
		for  {
			if len(e.RequestQ)>0 && len(e.WorkerQ)>0 {
				e.WorkerQ.pop().in <- e.RequestQ.pop()
			}else{
				select{
				case r := <- e.RequestChan :
					e.RequestQ.push(r...)
				case w := <-e.WorkerChan :
					e.WorkerQ.push(w)
				}
			}
		}
	}()

	// deal with Parser result
	go func() {
		n:=0
		for  {
			res := <- e.ParserResChan
			if len(res.Requests)>0 {
				e.RequestChan <- res.Requests
			}

			for _,item := range res.Items  {
				new := item.(model.Article)
				fmt.Printf("Got Item #%d: %v\n",n,new.Title)

				model.AddArticle(new)
				//item = item
				n+=1
			}
			//timer:=time.Tick(time.Second)
			//<-timer
		}
	}()

	timer := time.Tick(time.Second*5)
	n := 0

	file := "./" +"log"+ ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetFlags(log.LstdFlags)

	for  {
		var workersId []int
		for _,w := range e.WorkerQ{
			workersId = append(workersId,w.id)
		}

		log.Printf(" - running state #%d -------\nrequest queue length : %d , available worker count %d , parser result queue leng %d\n%v\n-------------------------------\n",n,len(e.RequestQ),len(e.WorkerQ),len(e.ParserResChan),workersId)
		//fmt.Println(workersStatus)
		n+=1
		<- timer
	}



}

type Worker struct {
	id int
	in chan Request
	out chan ParserRes
	status string
}

func (e *Concurrent) CreateWorker(id int,out *chan ParserRes) Worker {
	w:=Worker{
		id:id,
		in:make(chan Request),
		out:*out,
		status: "available",
	}
	go func() {
		for {
			// waiting for assignments
			request:= <- w.in
			//fmt.Println("got a task")

			// change worker ststus
			w.status = "working"

			// start to work
			parserRes, err := work(request)
			if err!=nil {
				log.Printf("wrong url %s : %s",request.Url,err)
				continue
			}

			//fmt.Println("send result")
			w.out <- parserRes  // send the results
			//fmt.Println("send result success")
			//e.WorkerQ.push(w)  // push into available workers queue
			//fmt.Println("talk i am free")
			e.WorkerChan <- w
			//fmt.Println("free")
		}
	}()

	return w
}

