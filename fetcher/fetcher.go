package fetcher

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func Fetch(url string) (io.Reader,error) {
	// 减缓请求速度
	rand.Seed(time.Now().Unix())
	speed:= rand.Intn( 800-200 ) + 800
	<- time.Tick(time.Millisecond* time.Duration(speed) )

	res,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	//defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code: %d",res.StatusCode)
	}

	return res.Body,nil
}