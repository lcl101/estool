package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/lcl101/estool/mylog"
	"github.com/lcl101/estool/search"
	_ "github.com/lcl101/estool/statik"
	"github.com/rakyll/statik/fs"
)

// Config ...
type Config struct {
	ServerHost string
}

var config *Config

func main() {
	config = &Config{}
	flag.StringVar(&config.ServerHost, "p", "0.0.0.0:9800", "服务端口")
	flag.Parse()
	statikFS, err := fs.New() // 静态文件编译成二进制

	if err != nil {
		// log.Fatalf(err.Error())
		// mylog.Error(err.Error())
		mylog.Error("静态文件错误，%v", err.Error())
	}
	// http.FileServer(statikFS) // 加载http server file

	var mux = http.NewServeMux()
	// router
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	// 加载search package
	search.NewSearch(search.InitOptions{
		Mux: mux,
	})

	// start http server
	go func(serverAddr string, m *http.ServeMux) {
		if err = http.ListenAndServe(serverAddr, m); err != nil {
			// fmt.Println("Cant start server:", err)
			mylog.Error("不能启动服务，%v", err)
			os.Exit(1)
		}
	}(config.ServerHost, mux)

	mylog.Debug("服务器启动: %s", config.ServerHost)
	// openPage()
	// 捕捉ctrl+C 退出信号
	handleSignals()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	mylog.Debug("停止服务")
}
