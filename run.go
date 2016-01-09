package Run

import (
	"sync"
)

type Run struct {
	in     chan interface{}              // 接受传入数据
	inover chan bool                     //接受数据完毕
	result chan interface{}              //返回产生结果
	wait   sync.WaitGroup                //处理多个Goroutine
	deal   func(interface{}) interface{} //处理函数
}

func (r *Run) In(item interface{}) {
	r.in <- item
}
func (r *Run) InOver() {
	go func() {
		r.inover <- true
	}()
}
func (r *Run) ChanIn(in chan interface{}) {
	go func() {
		for item := range in {
			r.in <- item
		}
		r.inover <- true
	}()
}

func (r *Run) Result() chan interface{} {
	return r.result
}
func (r *Run) CloseInChan() {
	close(r.in)
}

func (r *Run) CloseInOverChan() {
	close(r.inover)
}

func (r *Run) CloseResultChan() {
	close(r.result)
}

func (r *Run) Deal() {
	r.wait.Add(1)
	go func() {
		for item := range r.in {
			result := r.deal(item)
			if result != nil {
				r.result <- result
			}
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
	<-r.inover
}

func NewRun(DealFn func(interface{}) interface{}) *Run {
	return &Run{
		in:     make(chan interface{}),
		inover: make(chan bool, 1),
		result: make(chan interface{}),
		deal:   DealFn,
	}
}
