package main

import (
	"fmt"
	"time"
)

//	func main() {
//		//fmt.Println(time.Now().Unix())            //获取当前时间戳 单位是秒
//		//fmt.Println(time.Now().UnixNano())        //获取当前时间戳 单位是纳秒
//		//fmt.Println(time.Now().UnixMilli() / 1e6) //将纳秒转换为毫秒
//		//fmt.Println(time.Now().UnixMicro() / 1e9) //将纳秒转换为微秒
//		//
//		//fmt.Println(time.Unix(time.Now().UnixNano()/1e9, 0)) //将毫秒转换为 time 类型
//		//// 获取指定时间的时间戳。只要是time类型都能获得时间戳
//		//dt, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51")
//		//fmt.Println(dt.Unix())
//
// }
//
//	func main() {
//		a :=  time.Second
//		round := a.Round(time.Second)
//		fmt.Println(round)
//		fmt.Println(a.Round(10 * time.Second))
//		fmt.Println(a.Round(time.Minute))
//		fmt.Println(a.Round(time.Hour))
//		fmt.Println(a.Round(10 * time.Hour))
//	}
//
//	func main() {
//		now := time.Now().Hour()
//		fmt.Println(now)
//		hour, minutes, seconds := time.Now().Clock()
//		fmt.Println(hour, minutes, seconds)
//	}
func main() {

	// 定时器表示在未来某一时刻的独立事件。
	// 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。
	// 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(time.Second * 114)

	// `<-timer1.C` 直到这个定时器的通道 `C` 明确的发送了
	// 定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 如果你需要的仅仅是单纯的等待，你需要使用 `time.Sleep`。
	// 定时器有用的原因之一就是你可以在定时器失效之前，取消这个定时器。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
