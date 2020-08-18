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

// TODO: Return a list of humanized strings for each directory
func getDirSize(path string) (int64, error) {
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

	return size, err
}

func main() {
	arg.MustParse(&appArgs)

	if appArgs.DirName != "" {
		if appArgs.ListSubDirs {
			//listDirs()
		} else {
			dirSize, err := getDirSize(appArgs.DirName)

			if err != nil {
				fmt.Println("Error getting size of directory")
			} else {
				fmt.Println(humanize.Bytes(uint64(dirSize)))
			}
		}
	} else {
		fmt.Println("Using default")
	}

}
