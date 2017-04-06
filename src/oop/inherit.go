package oop
/*
 Create by sunyu ,thanks.
*/
import "fmt"
import "strings"

type Inherit struct{
	Msg string
}

func (i *Inherit) Init() {
	println(i.Msg)
	syn <- "Inherit"
}

func (s *Inherit) Test() {
	msg := <-syn //阻塞等待通道
	if strings.EqualFold(msg, "Inherit") {
		persons()
	}
}

type Person struct {
	name   string
	age    int
	weight int
}

type Skills  []string
type Student struct {
	Person //继承
	Skills
	int
	spe string
}

func persons(){
	xuxu := Student{Person{"xuxu", 25, 68}, []string{"anatomy"}, 1, "boy"} //方式一,全部指定
	fmt.Println(xuxu)
	jane := Student{Person:Person{"Jane", 25, 100}, spe:"Biology"}  //方式二,指哪打哪
	fmt.Println(jane)
}



