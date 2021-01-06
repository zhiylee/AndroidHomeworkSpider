package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	res,_ := http.Get("https://new.qq.com/rain/a/20201202A0FX3G00")


	doc,err:=goquery.NewDocumentFromReader(res.Body)
	if err!=nil {
		log.Print(err)
	}

	content := doc.Find(".content-article").Text()

	fmt.Println(content)
}
