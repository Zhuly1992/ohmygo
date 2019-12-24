package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var count int

//互斥锁，锁定互斥锁，不是锁代码，当代码执行到有锁的地方，获取不到互斥锁的锁定，就会阻塞，达到同步
//读写锁，读锁之间不锁定，读写之间锁定
func main() {
	//ch := make(chan struct{}, 2)
	//var a = 0
	//var lock sync.Mutex
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		lock.Lock()
	//
	//		defer lock.Unlock()
	//		a += 1
	//		//fmt.Printf("go routines %d,a=%d,\n", i, a)
	//		ch <- struct{}{}
	//	}()
	//}
	//for i := 0; i < 100; i++ {
	//	<-ch
	//}
	//fmt.Println("done")
	//
	//for i := 0; i < 5; i++ {
	//	go read(i, ch)
	//}
	//for i := 0; i < 5; i++ {
	//	go write(i, ch)
	//}
	//for i := 0; i < 10; i++ {
	//	<-ch
	//}

	ch := make(chan struct{}, 5)

	// 新建 cond
	var l sync.Mutex
	cond := sync.NewCond(&l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			// 条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}

			fmt.Printf("goroutine%d 执行结束\n", i)
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Millisecond * 20)
	// 锁定一下，我要改变 count 的值
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}

}
func read(n int, ch chan struct{}) {
	//rw.RLock()
	//fmt.Printf("goroutine %d 进入读操作...\n", n)
	//v := count
	//fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	//rw.RUnlock()
	//ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	//rw.Lock()
	//fmt.Printf("goroutine %d 进入写操作...\n", n)
	//v := rand.Intn(1000)
	//count = v
	//fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	//rw.Unlock()
	//ch <- struct{}{}
}
