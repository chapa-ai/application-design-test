package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	route "github.com/chapa-ai/application-gaspar/internal/http"
	"github.com/chapa-ai/application-gaspar/internal/util"
)

func main() {
	srv := route.GetServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			util.LogFatalf("listen:%+v", err)
		}
	}()

	util.LogInfo("server start listen %s", srv.Addr)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		util.LogFatalf("server shutdown failed:%+v", err)
	}
	util.LogInfo("server closed")
}
