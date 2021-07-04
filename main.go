package main

import (
	"fmt"

	"github.com/NunChatSpace/queuue_golang/model"
	"github.com/NunChatSpace/queuue_golang/queue"
)

func main() {
	dbwt := queue.GetDBWriter()
	go dbwt.Execute()

	for i := 0; i < 1000; i++ {
		tbl := model.Table{
			ID:      uint(i),
			Column1: fmt.Sprintf("Column1_%d", i),
			Column2: fmt.Sprintf("Column2_%d", i),
		}

		dbwt.EnQueue(tbl)
	}

	// time.Sleep(2 * time.Second)
	dbwt.StopExecute()
}
