package main

import (
	"runtime/debug"
	"time"

	"fmt"
	"runtime"

	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
)

// 可以用来分析本地运行的程序，开箱即用

func main() {
	go func() {
		viewer.SetConfiguration(viewer.WithMaxPoints(1000))
		mgr := statsview.New()
		// Start() runs a HTTP server at `localhost:18066` by default.
		mgr.Start()
		// Stop() will shutdown the http server gracefully
		// mgr.Stop()
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			data := make([]byte, 512)
			fmt.Println(string(data))
			runtime.GC()
			time.Sleep(20 * time.Millisecond)
		}()
	}

	for i := 0; i < 1000; i++ {
		debug.FreeOSMemory()
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println("exit.")
	debug.FreeOSMemory()
	// busy working....
	//time.Sleep(time.Minute)
	select {}
}
