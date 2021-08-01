
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g := new(errgroup.Group)
	g.Go(func() error {
		return startServer(":8080")
	})
	g.Go(func() error {
		return startServer(":8081")
	})
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT)

	for{
		s:= <-ch
		switch s {
			case syscall.SIGQUIT:
         //TODO stop server
			default:
				if err := g.Wait();err!=nil{
					fmt.Printf("%v", err)
				}
		}
	}

	time.Sleep(time.Second*5)

	ch<-syscall.SIGQUIT

}

func startServer(addr string, ctx context.Context) error {
	http.HandleFunc("/"+addr, func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"hello")
	})
	srv := &http.Server{Addr: addr}
	return srv.ListenAndServe()
}
