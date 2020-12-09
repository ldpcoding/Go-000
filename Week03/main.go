package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 用于控制关闭的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, _ := errgroup.WithContext(ctx)
	// 启动多个 http 服务
	eg.Go(func() error {
		return startServer("8080", ctx)
	})
	eg.Go(func() error {
		return startServer("8081", ctx)
	})
	// 监听关闭信号
	eg.Go(func() error {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)

		for {
			select {
			case s := <-sig:
				cancel()
				return errors.New("server shutdown by signal " + s.String())
			case <-ctx.Done():
				return nil
			}
		}
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func startServer(port string, ctx context.Context) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, port)
	})
	s := &http.Server{Addr: ":" + port, Handler: mux}
	s.RegisterOnShutdown(func() {
		fmt.Println("server " + port + " is shutting down")
	})
	go func() {
		<-ctx.Done()

		sdCtx, sdCancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer sdCancel()
		_ = s.Shutdown(sdCtx)
	}()
	fmt.Println("server " + port + " is starting")
	return s.ListenAndServe()
}
