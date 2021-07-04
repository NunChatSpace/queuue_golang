package queue

import (
	"fmt"

	"github.com/NunChatSpace/queuue_golang/model"
)

type IDBWriter interface {
	EnQueue(model.Table)
	Execute()
	StopExecute()
}

type DBWriter struct {
	Queue chan model.Table
	Break chan int
}

func GetDBWriter() IDBWriter {
	ch := make(chan model.Table)
	break_ch := make(chan int, 1)
	db := DBWriter{
		Queue: ch,
		Break: break_ch,
	}
	return db
}

func (dw DBWriter) EnQueue(t model.Table) {
	dw.Queue <- t
}

func (dw DBWriter) Execute() {
	for {
		select {
		case t := <-dw.Queue:
			fmt.Printf("Write :%v\n", t)
		case <-dw.Break:
			fmt.Println("Quit")
			return
		}
	}
}

func (dw DBWriter) StopExecute() {
	dw.Break <- 1
}
