package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/brettski/go-termtables"
	"github.com/dustin/go-humanize"
)

// Add total file sizes at the end of the table
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

	table := termtables.CreateTable()
	table.AddHeaders("Name", "Size")

	for _, file := range files {
		if file.IsDir() {
			dirSize := getDirSize(path.Join(dirPath, file.Name()))
			totalSize = totalSize + dirSize
			table.AddRow(file.Name(), getHumanizedSize(dirSize))
		}
	}

	table.AddSeparator()
	table.AddRow("TOTAL", getHumanizedSize(totalSize))
	fmt.Println(table.Render())
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
