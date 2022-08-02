package main

import (
	"fmt"
	"sync"
	"time"
)

func ChanSaleTicket(windowIndex int) int {
	Tickets--
	fmt.Println(windowIndex, "---saleTicket ----", Tickets)
	return Tickets
}

func ChanGetTicket(windowIndex int) int {
	fmt.Println(windowIndex, "---getTickets ----", Tickets)
	return Tickets
}

func ChanWindow(windowIndex int) {
	fmt.Println(windowIndex, "---wind start ----")
	time.Sleep(1 * time.Second) //这里是为了等待所有窗口(goroutine)都创建好
	for {
		// mutex.Lock()
		myChan <- windowIndex
		if ChanGetTicket(windowIndex) > 0 { //有余票继续售卖
			time.Sleep(1 * time.Second) //查到余票后先进行一些操作(让其同步更新 出现问题)
			fmt.Println(windowIndex, "sale success", ChanSaleTicket(windowIndex))
			// mutex.Unlock()
			<-myChan
		} else { //没有余票了 关门下班
			// mutex.Unlock()
			<-myChan
			break
		}
	}
	EndChan <- windowIndex
}

//票数
var Tickets int = 5

//窗口数
var Window int = 3

//锁
var myChan = make(chan int, 1)
var EndChan = make(chan int, Tickets)
var mutex sync.RWMutex

func main() {

	// var wait sync.WaitGroup
	// wait.Add(Window)
	for i := 0; i < Window; i++ {
		go ChanWindow(i)
	}

	// wait.Wait()
	i := 0
	for {
		select {
		case <-EndChan:
			i++
			fmt.Println("get one goroutine run over ", i, "chan len= ", len(EndChan))
			if i == Window {
				fmt.Println(fmt.Sprintf("all goroutine run over.  bye..."))
				close(EndChan)
				return
			}
		}
	}
}
