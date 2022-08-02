package main

import (
	"fmt"
	"sync"
	"time"
)

func mySaleTicket(windowIndex int) (int, int) {
	rwmutex.Lock() //售票时加上写锁  等待其他读锁释放
	defer rwmutex.Unlock()

	if Tickets <= 0 {
		return -1, Tickets

	}
	Tickets--
	fmt.Println(windowIndex, "---saleTicket ----", Tickets)
	return 1, Tickets
}

//这里关于查票的情况 其实不加读锁也应该是可以的，因为实际售票方法中加的是读锁，避免了超卖问题
func getTicket(windowIndex int) int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	fmt.Println(windowIndex, "---getTickets ----", Tickets)
	return Tickets
}

func window(windowIndex int, wait *sync.WaitGroup) {
	fmt.Println(windowIndex, "---wind start ----")
	time.Sleep(1 * time.Second) //这里是为了等待所有窗口(goroutine)都创建好
	for {
		if getTicket(windowIndex) > 0 { //有余票继续售卖
			time.Sleep(1 * time.Second) //查到余票后先进行一些操作(让其同步更新 出现问题)
			err, _ := mySaleTicket(windowIndex)
			fmt.Println(fmt.Sprintf("%d sale [status] %d", windowIndex, err))
		} else { //没有余票了 关门下班
			break
		}
	}
	wait.Done()
}

//票数
var Tickets int = 5

//窗口数
var Window int = 3

//锁
var rwmutex sync.RWMutex

func main() {

	var wait sync.WaitGroup
	wait.Add(Window)
	for i := 0; i < Window; i++ {
		go window(i, &wait)
	}

	wait.Wait()
	fmt.Println("end ticket -----", Tickets)
}
