package engine

import (
	"fmt"
	"log"
)

type Single struct{

}

func (e *Single) Run(seeds ...Request)  {
	var queue []Request

	for _,seed := range seeds {
		queue=append(queue,seed)
	}

	for len(queue)>0{
		//pop request
		request:= queue[0]
		queue = queue[1:]

		// print status
		log.Printf("carwling %s",request.Url)

		// work
		parserRes, err := work(request)
		if err != nil {
			log.Println(err)
			continue
		}

		queue = append(queue,parserRes.Requests...)  //feed requests into queue

		//print all item data
		for _,item := range parserRes.Items{
			fmt.Println("Got item : ",item)
		}


	}
}
