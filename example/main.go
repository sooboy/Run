package main

import (
	"github.com/sooboy/Run"
	"fmt"
)
func DealData(data interface{}) interface{} {
	return data
}

func main() {
	//实例化一个新的Run
	var run = Run.NewRun(DealData)
	//处理结果
	go func() {
		for {
			select {
			case item, ok := <-run.Result():
				if ok {
					fmt.Println(item)
				} else {
					return
				}
			}
		}
	}()
	//开启5个Goroutine并发处理
	run.DealWith(5)
	//发送10个数据
	for i := 0; i < 100; i++ {
		run.In(true)
	}

	run.Wait()

}
