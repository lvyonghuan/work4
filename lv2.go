package main

import "sync"

//前排：这次只完成了最基础的部分。只调了每半分钟的消息，整点消息还没调过，应该没bug吧（心虚
//感觉用time.Now().Clock()做还是不太合适
import (
	"fmt"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan int, 1)

func main() {
	for {
		wg.Add(2)
		go timeer()
		go printf()
		wg.Wait()
	}
}
func timeer() {
	hour, minute, second := time.Now().Clock()
	if second == 59 || second == 29 {
		time.Sleep(time.Second)
		ch <- 0
	} else if second == 0 {
		if minute == 0 {
			if hour == 2 {
				ch <- 1
			} else if hour == 6 {
				ch <- 2
			} else {
				ch <- -1
			}
		} else {
			ch <- -1
		}
	} else {
		ch <- -1
	}
	wg.Done()
}
func printf() {
	temp := <-ch
	if temp == 0 {
		fmt.Println("现在是幻想时间。")
	}
	if temp == 1 {
		fmt.Println("再来，再来一回合！")
	}
	if temp == 2 {
		fmt.Println("忽接神罗书一封，原来恶名传契丹")
	}
	wg.Done()
}
