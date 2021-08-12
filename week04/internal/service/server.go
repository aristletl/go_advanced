package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"week04/internal/conf"
	"week04/internal/server"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	s      http.Server
}

func New(c *conf.Data, ser *server.Server) *Server {
	var s Server
	s.s = http.Server{
		Addr:    c.Server.Addr,
		Handler: ser.GetMux(),
	}

	return &s
}

func (s *Server) Run() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {
		defer fmt.Println("Listen defer")
		return s.s.ListenAndServe()
	})

	eg.Go(func() error {
		<-ctx.Done()
		return s.s.Shutdown(s.ctx)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				s.cancel()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
