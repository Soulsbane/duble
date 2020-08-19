package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/brettski/go-termtables"
	"github.com/dustin/go-humanize"
)

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

func listDirs(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Failed to read directory")
		return
	}

	table := termtables.CreateTable()
	table.AddHeaders("Name", "Size")

	for _, file := range files {
		if file.IsDir() {
			table.AddRow(file.Name(), getHumanizedDirSize(file.Name()))
		}
	}

	fmt.Println(table.Render())
}

func getHumanizedDirSize(path string) string {
	dirSize := getDirSize(path)
	humanizedStr := humanize.Bytes(uint64(dirSize))

	return humanizedStr
}

func main() {
	var appArgs struct {
		DirName     string `arg:"positional"`
		ListSubDirs bool   `arg:"-l, --list" default:"false" help:"List diretories under the passed directory"`
	}

	arg.MustParse(&appArgs)

	if appArgs.DirName != "" {
		if appArgs.ListSubDirs {
			listDirs(appArgs.DirName)
		} else {
			fmt.Println(getHumanizedDirSize(appArgs.DirName))
		}
	} else {
		path, _ := os.Getwd()
		listDirs(path)
	}

}
