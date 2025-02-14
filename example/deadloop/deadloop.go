package main

import (
	"net/http"
	"time"

	"github.com/ttbug/holmes"
)

func init() {
	http.HandleFunc("/deadloop", deadloop)
	go http.ListenAndServe(":10003", nil)
}

func main() {
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath("/tmp"),
		holmes.WithCPUDump(10, 25, 80),
	)
	h.EnableCPUDump().Start()
	time.Sleep(time.Hour)
}

func deadloop(wr http.ResponseWriter, req *http.Request) {
	for i := 0; i < 4; i++ {
		for {
			time.Sleep(time.Millisecond)
		}
	}
}
