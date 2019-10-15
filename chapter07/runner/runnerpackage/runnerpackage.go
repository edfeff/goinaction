package runnerpackage

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner 在指定的之间内执行任务
//中断信号时结束这些任务
type Runner struct {
	//存放中断信号
	interrupt chan os.Signal
	//错误信号
	complete chan error
	//超时
	timeout <-chan time.Time
	tasks   []func(int)
}

var ErrTimeout = errors.New("timeout")
var ErrInterrupt = errors.New("interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//任务 为函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	//希望接受所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

//运行并检查中断
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//invoke task
		task(id)
	}
	return nil
}

//判断中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
