package Run

import (
	"sync"
)

type Run struct {
	in     chan interface{}              // 接受传入数据
	result chan interface{}              //返回产生结果
	wait   sync.WaitGroup                //处理多个Goroutine
	deal   func(interface{}) interface{} //处理函数
}

func (r *Run) In(item interface{}) {
	r.in <- item
}
func (r *Run) Result() chan interface{} {
	return r.result
}
func (r *Run) CloseInChan() {
	close(r.in)
}
func (r *Run) CloseResultChan() {
	close(r.result)
}
func (r *Run) Deal() {
	r.wait.Add(1)
	go func() {
		for item := range r.in {
			r.result <- r.deal(item)
		}
		r.wait.Done()
	}()
}
func (r *Run) DealWith(num int) {
	for i := 0; i < num; i++ {
		r.Deal()
	}
}
func (r *Run) Wait() {
	r.CloseInChan()
	r.wait.Wait()
	r.CloseResultChan()
}
func NewRun(DealFn func(interface{}) interface{}) *Run {
	return &Run{
		in:     make(chan interface{}),
		result: make(chan interface{}),
		deal:   DealFn,
	}
}

