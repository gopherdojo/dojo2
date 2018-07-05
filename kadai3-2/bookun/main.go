package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/gopherdojo/dojo2/kadai3-2/bookun/client"
)

const (
	// ExitCodeOK - 正常終了
	ExitCodeOK = iota
	// ExitCodeParseFlagError - フラグエラー
	ExitCodeParseFlagError
	// ExitCodeCreateHTTPClient - HTTPClient作成失敗
	ExitCodeCreateHTTPClient
	// ExitCodeErrorDownload - DL中のエラー
	ExitCodeErrorDownload
	// ExitCodeErrorCansel - キャンセルによる終了
	ExitCodeErrorCansel
)

// CLI - outStream, errStreamを持つ
type CLI struct {
	outStream, errStream io.Writer
}

// Run - 分割DLを開始. キャンセルを受け付ける
func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("bget", flag.ContinueOnError)
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	url := args[1]
	client, err := client.NewClient(url)
	if err != nil {
		return ExitCodeCreateHTTPClient
	}
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	defer func() {
		signal.Stop(ch)
		cancel()
	}()
	fmt.Fprintln(c.outStream, "Start downloading")
	go func() error {
		select {
		case <-ch:
			cancel()
			return nil
		case <-ctx.Done():
			if err = client.DeleteFiles(); err != nil {
				return err
			}
			return nil
		}
	}()
	err = client.Download(ctx)
	if err != nil {
		return ExitCodeErrorDownload
	}
	fmt.Fprintln(c.outStream, "Finish downloading")
	return ExitCodeOK
}

func main() {
	cli := &CLI{os.Stdout, os.Stderr}
	os.Exit(cli.Run(os.Args))
}
