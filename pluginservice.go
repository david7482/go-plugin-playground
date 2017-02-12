package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"plugin"
	"strings"
	"sync"
	"syscall"

	"github.com/david7482/go-plugin-playground/common/logger"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.NewLogger("request")

	plugin_so := strings.Split(strings.ToLower(r.URL.Path), "/")[1] + ".so"

	p, err := plugin.Open(plugin_so)
	if err != nil || p == nil {
		w.WriteHeader(http.StatusNotFound)
		log.WithFields(logger.Fields{"plugin": plugin_so, "err": err.Error()}).WARN("Open plugin fail")
		return
	}

	symbol, err := p.Lookup("ServeHTTP")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.WithFields(logger.Fields{"plugin": plugin_so, "err": err.Error()}).WARN("Lookup symbol fail")
		return
	}

	log.WithFields(logger.Fields{"plugin": plugin_so, "symbol": symbol}).INFO("Invoke plugin symbol")
	serve_http := symbol.(func(http.ResponseWriter, *http.Request))
	serve_http(w, r)
}

func main() {
	http.HandleFunc("/", httpHandler)

	log := logger.NewLogger("main")
	log.INFO("Start WebService ...")

	// Start HTTP server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := http.Server{}

	// Start web service
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		defer wait.Done()
		server.Serve(lis)
	}()

	// Monitor SIGINT and SIGTERM in signal_channel
	sig_close_channel := make(chan os.Signal)
	signal.Notify(sig_close_channel, syscall.SIGINT, syscall.SIGTERM)

	// Wait until we got system signal
	select {
	case <-sig_close_channel:
		log.INFO("Got signal to graceful shutdown")
		server.Shutdown(context.TODO())
	}
	wait.Wait()
}
