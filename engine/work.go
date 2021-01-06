package engine

import (
	"androidHomeworkSpider/fetcher"
)

func work(request Request) (ParserRes,error) {
	//run the fetcher
	r, err := fetcher.Fetch(request.Url)
	if err != nil {
		//log.Printf("wrong fetch url %s : %v", request.Url,err)
		return ParserRes{},err
	}

	// run the parser
	return request.ParserFunc(r,request)
}