package main

// 引入oop下文件
import "oop"
// go标准库包
import "time"

func main() {
	closure := oop.Closure{"begin ==闭包测试=="}
	inherit := oop.Inherit{"begin ==继承测试=="}
	goRoutines := oop.Thread{"begin ==并发测试=="}
	goChannel := oop.GoChannel{"begin ==通道测试=="}
	//map集合int->Test
	tests := map[int]oop.Test{}
	tests[0] = &closure
	tests[1] = &inherit
	tests[2] = &goRoutines
	tests[3] = &goChannel
	//每个测试结果独立进行互不影响
	var syn chan string = make(chan string)
	//_ 表示不使用key，这里只迭代value
	for _, t := range tests {
		go func(){
			go t.Init()
			t.Test()
			time.Sleep(time.Second*5)
			syn <- "done"
		}()
		println(<-syn)
	}
}

type Weekday int

const (
	//iota在const关键字出现时将被重置为0(const内部的第一行之前)，
	//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string { return days[d] }
