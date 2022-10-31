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
	hidden "github.com/tobychui/goHidden"
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
	humanizedStr := humanize.Bytes(uint64(size)) // SI size.
	//humanizedStr := humanize.IBytes(uint64(size)) // IEC size.
	return humanizedStr
}

func outputTable(dirs map[string]int64, totalSize int64) {
	dirDataTable := table.NewWriter()
	dirDataTable.SetOutputMirror(os.Stdout)
	dirDataTable.AppendHeader(table.Row{"Name", "Size"})

	if len(dirs) > 0 {
		for dirName, dirSize := range dirs {
			dirDataTable.AppendRow(table.Row{dirName, getHumanizedSize(dirSize)})
		}
	}

	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", getHumanizedSize(totalSize)})
	dirDataTable.SetStyle(table.StyleRounded)
	dirDataTable.Render()
}

func listDir(dirPath string, showHidden bool) {
	var dirs = map[string]int64{}
	var totalSize int64

	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Failed to read directory")
	}

	for _, file := range files {
		if !file.IsDir() {
			isHidden, _ := hidden.IsHidden(file.Name(), false)
			totalSize += file.Size()

			if isHidden {
				if showHidden {
					dirs[file.Name()] = file.Size()
				}
			} else {
				dirs[file.Name()] = file.Size()
			}
		}
	}

	outputTable(dirs, totalSize)
}

func listDirs(dirPath string, showHidden bool) {
	var dirs = map[string]int64{}
	var totalSize int64
	var rootDirSize int64

	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Failed to read directory")
		return
	}

	for _, file := range files {
		if file.IsDir() {
			isHidden, _ := hidden.IsHidden(file.Name(), false)
			dirSize := getDirSize(path.Join(dirPath, file.Name()))

			if isHidden {
				if showHidden {
					dirs[file.Name()] = dirSize
					totalSize += dirSize
				}
			} else {
				dirs[file.Name()] = dirSize
				totalSize += dirSize
			}
		} else {
			rootDirSize += file.Size()
			totalSize += file.Size()
		}
	}

	dirs["Root Directory"] = rootDirSize
	outputTable(dirs, totalSize)
}

func main() {
	var appArgs args

	arg.MustParse(&appArgs)
	path := appArgs.DirName

	if path == "" {
		path, _ = os.Getwd()
	}

	if appArgs.ListRootFilesOnly {
		listDir(path, appArgs.ListAll)
	} else {
		listDirs(path, appArgs.ListAll)
	}
}
