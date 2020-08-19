package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
)

var appArgs struct {
	DirName     string `arg:"positional"`
	ListSubDirs bool   `arg:"-l, --list" default:"false" help:"List diretories under the passed directory"`
}

func getDirSize(path string) int64 {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return err
	})

	if err != nil {
		return 0
	}

	return size
}

func main() {
	arg.MustParse(&appArgs)

	if appArgs.DirName != "" {
		if appArgs.ListSubDirs {
			//listDirs()
		} else {
			dirSize := getDirSize(appArgs.DirName)
			fmt.Println(humanize.Bytes(uint64(dirSize)))
		}
	} else {
		fmt.Println("Using default")
	}

}
