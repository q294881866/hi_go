package oop

/**go 语言通道测试**/
import (
	"fmt"
	"time"
	"strings"
)

// 来自google的例子
func ping(pings chan<- string, msg string) {  
    pings <- msg  
}  
  
func pong(pings <-chan string, pongs chan<- string) {  
    msg := <-pings  
    pongs <- msg  
}  
  
func pingpong() {  
    pings := make(chan string, 1)  
    pongs := make(chan string, 1)  
    ping(pings, "乒乓球打过去了，接着")  
    pong(pings, pongs)  
    fmt.Println(<-pongs)  
}  


type GoChannel struct {
	Msg string
}
func (gc *GoChannel) Test() {
	msg:=<- syn//阻塞取数据
	if strings.EqualFold(msg, "GoChannel"){
		pingpong()
	}
}
func (gc *GoChannel) Init() {
	println(gc.Msg)
	syn <- "GoChannel"
}

// go中通道默认是阻塞的，必须接收端和发送端都准备好，否则会死锁
func deadlock() {
	var c1 chan string = make(chan string)
	func() {
		time.Sleep(time.Second * 2)
		c1 <- "res 1"
	}()

	println("c1 :", <-c1)
}
// 通道配合goroutines运行
func fixDeadlock1() {
	var c1 chan string = make(chan string)
	go func() {
		c1 <- "res 1"
	}()

	println("c1 :", <-c1)
}

// 通过channel加buffer
func fixDeadlock2() {
	var c1 chan string = make(chan string, 1)
	func() {
		//延时执行，模拟超时
		//time.Sleep(time.Second * 2)
		c1 <- "res 1"
	}()
	select {
	case msg := <-c1:
		fmt.Println("received message", msg)
	case <-time.After(time.Second):
		println("超时")
	}
}

