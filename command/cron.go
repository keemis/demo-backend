package command

import (
	"time"
)

// Run 任务
func Run() {

	go func() {
		for {
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	}()
}
