package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gopherdojo/dojo2/kadai3-2/pdownload"
)

var (
	pCountOpt    = flag.Int("p", 5, "分割数")
	outputDirOpt = flag.String("o", ".", "ダウンロードファイルの出力先ディレクトリ")
	tmpDirOpt    = flag.String("t", "/tmp/kaznishi_pdownload", "分割ファイルの一時格納ディレクトリ")
)

func main() {
	option := pdownload.Option{}
	option.Init()

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "ダウンロード対象のURLが指定されていません")
		return
	}
	option.TargetURL = flag.Args()[0]
	option.PCount = *pCountOpt
	option.OutputDir = *outputDirOpt
	option.TmpDir = *tmpDirOpt

	ctx, cancel := context.WithCancel(context.Background())

	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan int, 1)
	signal.Notify(sigCh, trapSignals...)

	var wgMain sync.WaitGroup
	go func() {
		wgMain.Add(1)
		if err := pdownload.Run(ctx, doneCh, option); err != nil {
			fmt.Println(err)
		}
		wgMain.Done()
	}()
	select {
	case sig := <-sigCh:
		cancel()
		wgMain.Wait()
		fmt.Println("Got signal", sig)
	case code := <-doneCh:
		if code == 0 {
			fmt.Println("Done!!!!!")
		} else {
			fmt.Println("Failed...")
		}
	}
}
