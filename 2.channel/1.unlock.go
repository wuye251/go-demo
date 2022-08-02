package main

import (
	"fmt"
	"sync"
	"time"
)

func saleTicket(windowIndex int) int {
	Tickets--
	fmt.Println(windowIndex, "---saleTicket ----", Tickets)
	return Tickets
}

func getTicket(windowIndex int) int {
	fmt.Println(windowIndex, "---getTickets ----", Tickets)
	return Tickets
}

func window(windowIndex int, wait *sync.WaitGroup) {
	fmt.Println(windowIndex, "---wind start ----")
	time.Sleep(1 * time.Second) //这里是为了等待所有窗口(goroutine)都创建好
	for {
		if getTicket(windowIndex) > 0 { //有余票继续售卖
			time.Sleep(1 * time.Second) //查到余票后先进行一些操作(让其同步更新 出现问题)
			fmt.Println(windowIndex, "sale success", saleTicket(windowIndex))
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

func main() {
	var wait sync.WaitGroup
	wait.Add(Window)
	for i := 0; i < Window; i++ {
		go window(i, &wait)
	}
	wait.Wait()
	fmt.Println("end ticket -----", Tickets)
}
