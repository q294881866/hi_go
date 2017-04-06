package oop

/**闭包测试，这个程序返回一个自增函数**/
import (
	"fmt"
	"strings"
)

type Closure struct {
	Msg string
}

func (c *Closure) Init() {
	println(c.Msg)
	syn <- "Closure"
}
func (c *Closure) Test() {
	msg:=<-syn//阻塞等待通道
	if strings.EqualFold(msg, "Closure"){
		nextInts()
	}
}

// 这里返回一个函数
func increment() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func nextInts(){
	nextInt := increment()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

