package oop
// package 类似命名空间，如oop.Test
/**
	此文件定义一个测试接口
	1.要求每个文件的测试类都要实现接口
	2.方便统一测试
**/
type Test interface {
	// 调用每个文件的test()完成测试
	Test()
	// 测试前的一些工作
	Init()
}

// 保证 Init方法在Test方法前执行
// go中通道默认是阻塞的，必须接收端和发送端都准备好，否则会死锁
var syn chan string = make(chan string)

