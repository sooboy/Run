package main

import (
	"github.com/sooboy/Run"
	"fmt"
	"time"
)

func DealData(data interface{}) interface{} {
	//fmt.Println(data, "第一层处理")
	return data
}

func main() {
	start := time.Now()
	//实例化一个新的Run
	var run = Run.NewRun(DealData)

	//结果收集
	var dealresult = func(data interface{}) interface{} {
		//	fmt.Println(data, "deal 2")
		return nil
	}

	var result = Run.NewRun(dealresult)
	result.DealWith(5)
	result.ChanIn(run.Result())

	//开启5个Goroutine并发处理
	run.DealWith(5)
	//发送10个数据

	for i := 0; i < 10000000; i++ {
		run.In("good")
	}

	run.InOver()
	run.Wait()
	//fmt.Println("结果接受完毕，处理第三层！")
	result.Wait()
	fmt.Println(time.Now().Sub(start))
}
