package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
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

func getHumanizedSize(size int64) string {
	humanizedStr := humanize.Bytes(uint64(size))
	return humanizedStr
}

func listDirs(dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	var totalSize int64

	if err != nil {
		fmt.Println("Failed to read directory")
		return
	}

	dirDataTable := table.NewWriter()
	dirDataTable.SetOutputMirror(os.Stdout)
	dirDataTable.AppendHeader(table.Row{"Name", "Size"})

	for _, file := range files {
		if file.IsDir() {
			dirSize := getDirSize(path.Join(dirPath, file.Name()))
			totalSize = totalSize + dirSize
			dirDataTable.AppendRow(table.Row{file.Name(), getHumanizedSize(dirSize)})
		}
	}

	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", getHumanizedSize(totalSize)})
	dirDataTable.Render()
}

func main() {
	var appArgs struct {
		DirName     string `arg:"positional"`
		ListSubDirs bool   `arg:"-l, --list" default:"false" help:"List directories under the passed directory."`
	}

	arg.MustParse(&appArgs)

	if appArgs.DirName != "" {
		if appArgs.ListSubDirs {
			listDirs(appArgs.DirName)
		} else {
			dirSize := getDirSize(appArgs.DirName)
			fmt.Println(getHumanizedSize(dirSize))
		}
	} else {
		path, _ := os.Getwd()
		listDirs(path)
	}
}
