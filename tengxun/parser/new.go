package parser

import (
	"androidHomeworkSpider/engine"
	"androidHomeworkSpider/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
)

func New(r io.Reader,request engine.Request) (engine.ParserRes,error){

	r = transform.NewReader(r,simplifiedchinese.GBK.NewDecoder())

	res := engine.ParserRes{}

	doc,err:=goquery.NewDocumentFromReader(r)
	if err!=nil {
		log.Print(err)
		return engine.ParserRes{},err
	}

	new := request.Data.(model.Article)

	new.Content,err = doc.Find(".content-article").Html()

	if err!=nil {
		log.Print(err)
		return engine.ParserRes{},err
	}

	//请求页面失败 被限流 重试处理
	if new.Content==""{
		log.Println("Warry : 抓取 ",new.Title," 失败，文章搬家了，舍弃")
		return engine.ParserRes{},nil
		//log.Println(new)
		//log.Println(request)
		//res.Requests = append(res.Requests,request)
		//
		//return res,nil
	}


	content := doc.Find(".content-article")

	content.Find("img").Each(func(i int, s *goquery.Selection) {
		imgSrc,exist:=s.Attr("src")
		if exist {
			s.SetAttr("src","http:"+imgSrc)
		}

		// 添加文章样式
		s.SetAttr("style","max-width: 100%;margin-top: 5px;")
	})

	contentHtml,_:=content.Html()

	m:=minify.New()
	m.Add("text/html",&html.Minifier{
		KeepEndTags: true,
		KeepQuotes: true,
	})
	new.Content,err = m.String("text/html",contentHtml)

	if err!=nil {
		log.Print(err)
		return engine.ParserRes{},err
	}

	res.Items = append(res.Items,new)

	return res,nil
}