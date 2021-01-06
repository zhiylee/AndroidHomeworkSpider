package parser

import (
	"androidHomeworkSpider/engine"
	"androidHomeworkSpider/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"time"
)

func NewList(r io.Reader,request engine.Request) (engine.ParserRes,error) {
	res := engine.ParserRes{}


	jsonStr,err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
		return res,err
	}

	var news map[string]interface{}

	json.Unmarshal( []byte(jsonStr), &news)

	newsData,ok := news["data"].(map[string]interface{})
	if !ok {
		log.Println("页面解析失败 没有下一页")
		return res,nil
	}
	newsList :=newsData["list"].([]interface{})

	for _,item := range newsList {
		new:=item.(map[string]interface{})

		// 检测文章是否存在
		isExist := model.ArticleIsExist(model.Article{
			Title: new["title"].(string),
		})
		if isExist {
			log.Println("Warry: 文章 ",new["title"].(string)," 已存在 跳过")
			continue
		}

		res.Requests=append(res.Requests,engine.Request{
			Url: new["url"].(string),
			ParserFunc: New,
			Data:model.Article{
				Title:        new["title"].(string),
				Cover:        new["img"].(string),
				CategoryId: request.Data.(map[string]interface{})["categoryId"].(int),
				CreateTime:   strTimeToUnix( new["create_time"].(string) ),
				CommentCount: 0,
				Author: new["media_name"].(string),
			},
		})

	}

	// 下一页
	u, err := url.Parse(request.Url)
	if err != nil {
		log.Println("url解析错误，下一页失败")
	}
	// 生成下一页url
	q:=u.Query()
	offset:=strToInt( q.Get("offset") )
	q.Set("offset", strconv.Itoa( offset + strToInt(q.Get("limit")) ) )

	u.RawQuery = q.Encode()
	log.Println("下一页加入任务队列：",u.String())
	// 生成Request
	res.Requests = append(res.Requests,engine.Request{
		Url:        u.String(),
		ParserFunc: NewList,
		Data:       request.Data,
	})

	return res,nil
}

func strTimeToUnix(from string) int {
	loc,_:=time.LoadLocation("Local")
	theTime,_ :=time.ParseInLocation("2006-01-02 15:04:05",from,loc)

	return int( theTime.Unix() )
}

func strToInt(s string) int {
	v,err:=strconv.ParseInt(s,10,0)
	if err != nil {
		return 0
	}

	return int(v)
}