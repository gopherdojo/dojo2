package main

import (
	"context"
	"dojo2/kadai3-2/mukai/loader"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var url *string
var split *int

func init() {
	url = flag.String("u", "", "download url")
	split = flag.Int("s", 10, "split size")
}

func main() {
	flag.Parse()

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
	if err := loader.Download(ctx, *url, *split); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
