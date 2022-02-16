package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/saracen/walker"
)

func getDirSize(path string) int64 {
	var size int64

	walkFn := func(path string, info os.FileInfo) error {
		size += info.Size()
		return nil
	}

	errorCallbackOption := walker.WithErrorCallback(func(pathname string, err error) error {
		if os.IsPermission(err) {
			return nil // INFO: Ignore permission errors
		}

		return err // INFO: Stop on all other errors
	})

	walker.Walk(path, walkFn, errorCallbackOption)

	return size
}

func getHumanizedSize(size int64) string {
	humanizedStr := humanize.Bytes(uint64(size))
	return humanizedStr
}

func listDir(dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	var dirSize int64

	if err != nil {
		fmt.Println("Failed to read directory")
	}

	dirDataTable := table.NewWriter()
	dirDataTable.SetOutputMirror(os.Stdout)
	dirDataTable.AppendHeader(table.Row{"Name", "Size"})

	for _, file := range files {
		if !file.IsDir() {
			dirSize += file.Size()
			dirDataTable.AppendRow(table.Row{file.Name(), getHumanizedSize(file.Size())})
		}
	}

	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", getHumanizedSize(dirSize)})
	dirDataTable.Render()
}

func listDirs(dirPath string) { // Maybe option for directories only
	files, err := ioutil.ReadDir(dirPath)
	var totalSize int64
	var rootDirSize int64

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
		} else {
			rootDirSize = rootDirSize + file.Size()
		}
	}

	dirDataTable.AppendRow(table.Row{"Root Directory", getHumanizedSize(rootDirSize)})
	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", getHumanizedSize(totalSize)})
	dirDataTable.Render()
}

func main() {
	var appArgs args
	path := appArgs.DirName

	if path == "" {
		path, _ = os.Getwd()
	}

	arg.MustParse(&appArgs)

	if appArgs.ListSubDirs {
		fmt.Println("Using --list version")
		listDirs(path)
	} else {
		listDir(path)
	}
}
