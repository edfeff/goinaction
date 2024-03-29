package work

import "sync"

//无缓存的通道
//工作接口
type Worker interface {
	Task()
}

//工作池
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutine)
	for i := 0; i < maxGoroutine; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//添加worker
func (p *Pool) Run(w Worker) {
	p.work <- w
}
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
