package main

import (
	"net/http"
	"time"

	"github.com/ttbug/holmes"
)

func init() {
	http.HandleFunc("/cpuex", cpuex)
	go http.ListenAndServe(":10003", nil)
}

func main() {
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath("/tmp"),
		holmes.WithCPUDump(20, 25, 80),
	)
	h.EnableCPUDump().Start()
	time.Sleep(time.Hour)
}

func cpuex(wr http.ResponseWriter, req *http.Request) {
	go func() {
		for {
			time.Sleep(time.Millisecond)
		}
	}()
}
