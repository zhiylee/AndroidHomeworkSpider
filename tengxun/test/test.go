package main

import (
	"androidHomeworkSpider/fetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"log"
	"strconv"
)

func main()  {
	r,_:=fetcher.Fetch("https://new.qq.com/omn/20201224/20201224A0D78Z00.html")

	r = transform.NewReader(r,simplifiedchinese.GBK.NewDecoder())

	doc,err:=goquery.NewDocumentFromReader(r)
	if err!=nil {
		log.Print(err)
	}



	content := doc.Find(".content-article")

	content.Find("img").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Attr("src"))
		imgSrc,exist:=s.Attr("src")
		if exist {
			s.SetAttr("src","http:"+imgSrc)
		}

		s.SetAttr("style","max-width: 100%;margin-top: 5px;")
	})

	contentHtml,_:=content.Html()

	m:=minify.New()
	m.Add("text/html",&html.Minifier{
		KeepEndTags: true,
		KeepQuotes: true,
	})
	b,_ := m.String("text/html",contentHtml)
	fmt.Println(b)

	//fmt.Println()

	//reg:=regexp.MustCompile(`<img.+src="(.+)">`)
	//match:=reg.FindAll([]byte(b),1)
	//for _,k:=range match{
	//	fmt.Println(string(k))
	//}

	//content,err:=goquery.NewDocumentFromReader(strings.NewReader(b))
	//if err!=nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(content.Html())


	//res,_ := http.Get("https://new.qq.com/rain/a/20201205A09JJJ00")
	//
	//
	//doc,err:=goquery.NewDocumentFromReader(res.Body)
	//if err!=nil {
	//	log.Print(err)
	//}
	//
	//content,_ := doc.Find(".content-article").Html()
	//
	//fmt.Println(content)


	//rand.Seed(time.Now().Unix())
	//
	//fmt.Println(time.Now().Unix())
	//<- time.Tick(time.Millisecond*2000)
	//fmt.Println(time.Now().Unix())

	// 下一页
	//u, err := url.Parse("https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=antip&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}")
	//if err != nil {
	//	log.Println("url解析错误，下一页失败")
	//}
	//// 生成下一页url
	//q:=u.Query()
	//fmt.Println(q.Get("offset"))
	//offset:=strToInt( q.Get("offset") )
	//fmt.Println(offset)
	//
	//fmt.Println( q.Get("limit") )
	//fmt.Println( strToInt( q.Get("limit") ) )
	//
	////q.Set("offset", string( offset + strToInt(q.Get("limit")) ) )
	//fmt.Println( offset + strToInt(q.Get("limit")) )
	//
	//a:= "30.98"
	//fmt.Println( strconv.Atoi(a) )
	//
	//
	//
	//u.RawQuery = q.Encode()
	//log.Println("下一页加入任务队列：",u.String())
	// 生成Request


}

func strToInt(s string) int {
	v,err:=strconv.ParseInt(s,10,0)
	if err != nil {
		return 0
	}

	return int(v)
}