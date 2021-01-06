package main

import (
	"androidHomeworkSpider/engine"
	"androidHomeworkSpider/tengxun/parser"
)

func main() {
	//e:=engine.Single{}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/",
	//	ParserFunc: parser.CityList,
	//})

	e:=engine.Concurrent{
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=antip&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
		ParserFunc: parser.NewList,
		Data: map[string]interface{}{
			"categoryId":1,
		},
	},
	engine.Request{
		Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=tech&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
		ParserFunc: parser.NewList,
		Data: map[string]interface{}{
			"categoryId":2,
		},
	},
		engine.Request{
			Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=ent&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
			ParserFunc: parser.NewList,
			Data: map[string]interface{}{
				"categoryId":3,
			},
		},
		engine.Request{
			Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=milite&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
			ParserFunc: parser.NewList,
			Data: map[string]interface{}{
				"categoryId":4,
			},
		},
		engine.Request{
			Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=world&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
			ParserFunc: parser.NewList,
			Data: map[string]interface{}{
				"categoryId":5,
			},
		},
		engine.Request{
			Url: "https://i.news.qq.com/trpc.qqnews_web.kv_srv.kv_srv_http_proxy/list?sub_srv_id=finance&srv_id=pc&offset=0&limit=20&strategy=1&ext={%22pool%22:[%22high%22,%22top%22],%22is_filter%22:10,%22check_type%22:true}",
			ParserFunc: parser.NewList,
			Data: map[string]interface{}{
				"categoryId":6,
			},
		})

}