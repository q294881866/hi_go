package oop

/**go 语言并发相关 这里命名参照java，并不是和恰当 **/
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"strings"
)

type Thread struct {
	Msg string
}

func (t *Thread) Test() {
	msg := <-syn //阻塞等待通道
	if strings.EqualFold(msg, "Thread") {
		concurrent()
		conWriteOrRead()
		mutexes()
	}
}
func (t *Thread) Init() {
	println(t.Msg)
	syn <- "Thread"
}

// =========================================go并发编程================================================
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func mutexes() {
	// 读写次数
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// The `reads` and `writes` channels
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// state map中读写，
	// read.resp  读结果保存如通道
	// write.resp 是否写入通道
	go func() {
		var state = make(map[int]int)
		for { //死循环
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 100个并发读入reads中
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10个并发写入writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 等待结果
	time.Sleep(time.Second)

	// 数据统计
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}


// =========================================go并发编程================================================

// 并发库
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Millisecond)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func concurrent() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		//开启并发任务 go 关键字
		go worker(w, jobs, results)
	}

	//发送
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	//发送完成
	close(jobs)

	// 接收消息
	for a := 1; a <= 5; a++ {
		<-results
	}
}

func conWriteOrRead() {
	// 保证并发map集合的安全
	var state = make(map[int]int)

	// 互斥锁，这里用于state
	var mutex = &sync.Mutex{}

	// 记录并发读写次数，用原子操作库
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// 100个线程重复读
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// [0-5)随机访问
				key := rand.Intn(5)
				// 同步控制
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				// 原子操作
				atomic.AddUint64(&readOps, 1)
				// 等待读
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 同样开10个并发任务对于写
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 等待10个并发写结果
	time.Sleep(time.Second)

	// 数据统计
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// 集合数据集
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
