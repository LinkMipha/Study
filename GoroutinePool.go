package main

import (
	"fmt"
	"time"
)

type Task struct {
	Taskid int
	f func()error
}

func TaskInit(taskid int,st func()error)*Task  {
	return &Task{
		Taskid: taskid,
		f: st,
	}
}

//运行
func (t *Task)Execute()  {
	t.f()
}

type Pool struct {
	TaskNum int
	EntryChan chan *Task
	WorkChan chan *Task
}

func NewPool(num int) *Pool {
	return  &Pool{
		TaskNum:   num,
		EntryChan: make(chan *Task),
		WorkChan:  make(chan *Task),
	}
}

func (p *Pool)Worker(id int)  {
	for task:=range p.WorkChan{
		task.Execute()
		fmt.Println("work id",id,"task id",task.Taskid,"is done")
	}
}

func (p *Pool)PoolRun()  {
	for i:=0;i<p.TaskNum;i++{
		go p.Worker(i)
	}
	for task:=range p.EntryChan{
		p.WorkChan <-task
	}
}

func task()error {
	fmt.Println(time.Now(), "Do something")
	return nil
}

/*
func main()  {
	p := NewPool(3)

	id :=0
	go func() {
		for {
			p.EntryChan <- TaskInit(id,task)
			id++
		}
	}()

	p.PoolRun()

}*/