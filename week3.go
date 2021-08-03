package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	g, Ctx :=errgroup.WithContext(context.Background())

	srv1 := &http.Server{Addr: ":8080"}
	srv2 := &http.Server{Addr: ":8081"}

	g.Go(func() error {
		return srv1.ListenAndServe()
	})
	g.Go(func() error {
		return srv2.ListenAndServe()
	})


	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGQUIT)

	stopServer := func() {
		srv1.Shutdown(Ctx)
		srv2.Shutdown(Ctx)
	}

	g.Go(func() error {
		s := <-signals
		switch s {
		case syscall.SIGQUIT:
			stopServer()
		}
		return errors.New("SIGQUIT")
	})

	g.Go(func() error {
		<-Ctx.Done()
		stopServer()
		return errors.New("SERVER ERROR")
	})

	time.Sleep(time.Second*5)
	//模拟退出信号
	signals<-syscall.SIGQUIT

	if err := g.Wait();err!=nil{
		fmt.Println(err)
	}
}
