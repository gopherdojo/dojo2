package main

import (
	"os"
	"strings"
	"testing"
)

var cli = &CLI{os.Stdout, os.Stdout}

func TestRunFlagParse(t *testing.T) {
	t.Run("Flag option error", func(t *testing.T) {
		args := strings.Split("./convert-cli -m hoge test_images", " ")
		errNum := cli.Run(args)
		if errNum != ExitCodeParseFlagError {
			t.Error(errNum)
		}
	})
	t.Run("Unspecified directory", func(t *testing.T) {
		args := strings.Split("./convert-cli", " ")
		errNum := cli.Run(args)
		if errNum != ExitCodeParseFlagError {
			t.Error(errNum)
		}
	})
}

func TestCreateFormat(t *testing.T) {
	args := strings.Split("./convert-cli -s testS -d testD test_images", " ")
	errNum := cli.Run(args)
	if errNum != ExitCodeCreateFormat {
		t.Error(errNum)
	}

}

func TestRunDirectory(t *testing.T) {
	args := strings.Split("./convert-cli hogeDir", " ")
	errNum := cli.Run(args)
	if errNum != ExitCodeSearchError {
		t.Error(errNum)
	}
}

func TestRun(t *testing.T) {
	argsPatterns := []string{
		"./convert-cli test_images",
		"./convert-cli -s jpeg -d png test_images",
		"./convert-cli -s png -d jpeg test_images",
	}
	for _, argsPattern := range argsPatterns {
		args := strings.Split(argsPattern, " ")
		code := cli.Run(args)
		if code != ExitCodeOK {
			t.Error(code)
		}
	}
}
