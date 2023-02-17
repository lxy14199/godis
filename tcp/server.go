package tcp

import (
	"context"
	"fmt"
	"godis/interface/tcp"
	"godis/lib/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("bind: %s, start listener...", cfg.Address))
	ListenAndServe(listener, handler, closeChan)
	return nil
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	errCh := make(chan error)
	defer close(errCh)
	go func() {
		select {
		case <-closeChan:
			logger.Info("get exit signal")
		case er := <-errCh:
			logger.Info(fmt.Sprintf("accpt error: %s", er.Error()))
		}
		logger.Info("shutting down...")
		_ = listener.Close()
		_ = handler.Close()
	}()

	ctx := context.Background()
	var waiDone sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			errCh <- err
			break
		}
		logger.Info("accpt link")
		waiDone.Add(1)
		go func() {
			defer func() {
				waiDone.Done()
			}()
			handler.Handle(ctx, conn)
		}()
	}
	waiDone.Wait()
}
