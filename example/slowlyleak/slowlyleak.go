package main

import (
	"net/http"
	"time"

	"github.com/ttbug/holmes"
)

func init() {
	http.HandleFunc("/leak", leak)
	go http.ListenAndServe(":10003", nil)
}

func main() {
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath("/tmp"),
		holmes.WithTextDump(),
		holmes.WithGoroutineDump(10, 25, 80),
	)
	h.EnableGoroutineDump().Start()
	time.Sleep(time.Hour)
}

func leak(wr http.ResponseWriter, req *http.Request) {
	taskChan := make(chan int)
	consumer := func() {
		for task := range taskChan {
			_ = task // do some tasks
		}
	}

	producer := func() {
		for i := 0; i < 10; i++ {
			taskChan <- i // generate some tasks
		}
		// forget to close the taskChan here
	}

	go consumer()
	go producer()
}
