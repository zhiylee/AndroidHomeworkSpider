package engine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestNilParser2(t *testing.T) {
	for i:=0;i<10;i++{
		go func() {
			timer:=time.Tick(time.Second)
			for  {
				<- timer
				//runtime.Gosched()
				
			}
		}()
	}
	
	go func() {
		i:=0
		for  {
			if i+=1;i%10==0 {
				// do someting
				a:=989898*4334343
				a+=1
			}else{
				runtime.Gosched()
			}
			//timer:=time.Tick(time.Microsecond)
			//<- timer
		}
	}()

	timer := time.Tick(time.Second)
	for  {
		<- timer
		//time.Sleep(time.Second)
		fmt.Println(time.Now().UnixNano())
	}
}

func TestNilParser3(t *testing.T) {
	ch:=make(chan int)


	go func() {
		for  {
			fmt.Println(<-ch)
		}
	}()

	ch <- 0
	ch <- 1
	timer:=time.Tick(time.Second)
	<-timer

}