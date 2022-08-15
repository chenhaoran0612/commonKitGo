package tools

import (
	"context"
	"net/http"
	"os"
	"strconv"
)

func GetHostname()(string, error){
	return os.Hostname()
}


var DEBUG_SERVER *http.Server
// 开启调试http服务
func StartDebugHttp(port int) error {
	// brew install graphviz
	// go tool pprof -http=:8080  'http://127.0.0.1:9999/debug/pprof/profile?seconds=60'
	// http.ListenAndServe(":9999", nil)

	addr := ":" + strconv.Itoa(
		port)
	DEBUG_SERVER = &http.Server{Addr: addr, Handler: nil}
	return DEBUG_SERVER.ListenAndServe()
}

// 关闭调试http
func ClosedDebugHttp() error {
	return DEBUG_SERVER.Shutdown(context.Background())
}