package main

import (
	"context"
	"dojo2/kadai3-2/mukai/loader"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signalCh := make(chan os.Signal, 1)
	signal.Notify(
		signalCh,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	go func() {
		<-signalCh
		cancel()
	}()
	go func() {
		<-ctx.Done()
		os.Exit(1)
	}()
	const Split = 200
	const Url = "https://www.noao.edu/image_gallery/images/d7/cygloop-8000.jpg"
	if err := loader.Download(ctx, Url, Split); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
