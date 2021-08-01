package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
)

type Serve struct {
	mux  *http.ServeMux
	stop chan struct{}
}

func (s *Serve) HandleFunc(pattern string, handler http.HandlerFunc) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.HandleFunc(pattern, handler)
}

func (s *Serve) BindServe() *Serve {
	return &Serve{stop: s.stop}
}

func (s *Serve) ShutDown() {
	close(s.stop)
}

func getUserName(rsp http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.FormValue("userName")
	str := fmt.Sprintf("Get your name is %s", userName)
	rsp.Write([]byte(str))
}

func serve(addr string, sev *Serve) error {
	s := http.Server{
		Addr:    addr,
		Handler: sev.mux,
	}
	go func() {
		<-sev.stop
		s.Shutdown(context.Background())
	}()
	fmt.Println("server listen:", addr)
	return s.ListenAndServe()
}

func registrySignal(s *Serve) error {
	signalCh := make(chan os.Signal, 1)
	fmt.Println("registry system signal......")
	signal.Notify(signalCh, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalCh
	s.ShutDown()
	return errors.Errorf("receive system signal: %s", sig.String())
}

func main() {
	stop := make(chan struct{})
	userServe := &Serve{
		stop: stop,
	}
	var g errgroup.Group
	g.Go(func() error {
		userServe.HandleFunc("/userName", getUserName)
		return serve(":8080", userServe)
	})
	g.Go(func() error {
		return registrySignal(userServe)
	})
	if err := g.Wait(); err != nil {
		fmt.Println("server broken with error: ", err)
	}
}
