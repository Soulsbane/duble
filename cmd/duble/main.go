package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/saracen/walker"
	hidden "github.com/tobychui/goHidden"
)

type DirInfo struct {
	Name string
	Size int64
}

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

func outputTable(dirs []DirInfo, totalSize int64) {
	dirDataTable := table.NewWriter()
	dirDataTable.SetOutputMirror(os.Stdout)
	dirDataTable.AppendHeader(table.Row{"Name", "Size"})

	if len(dirs) > 0 {
		for _, dirInfo := range dirs {
			dirDataTable.AppendRow(table.Row{dirInfo.Name, getHumanizedSize(dirInfo.Size)})
		}
	}

	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", getHumanizedSize(totalSize)})
	dirDataTable.SetStyle(table.StyleRounded)
	dirDataTable.Render()
}

func getListOfFiles(dirPath string, showHidden bool) ([]DirInfo, int64) {
	var dirs []DirInfo
	var totalSize int64

	files, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Failed to read directory")
	}

	for _, file := range files {
		info, _ := file.Info()

		if !file.IsDir() {
			isHidden, _ := hidden.IsHidden(file.Name(), false)
			totalSize += info.Size()

			if isHidden {
				if showHidden {
					dirInfo := DirInfo{file.Name(), info.Size()}
					dirs = append(dirs, dirInfo)
				}
			} else {
				dirInfo := DirInfo{file.Name(), info.Size()}
				dirs = append(dirs, dirInfo)
			}
		}
	}

	return dirs, totalSize
}

func getListOfDirs(dirPath string, showHidden bool) ([]DirInfo, int64) {
	var dirs []DirInfo
	var totalSize int64
	var rootDirSize int64

	files, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Failed to read directory")
	}

	for _, file := range files {
		info, _ := file.Info()

		if file.IsDir() {
			isHidden, _ := hidden.IsHidden(file.Name(), false)
			dirSize := getDirSize(path.Join(dirPath, file.Name()))

			if isHidden {
				if showHidden {
					dirInfo := DirInfo{file.Name(), dirSize}

					dirs = append(dirs, dirInfo)
					totalSize += dirSize
				}
			} else {
				dirInfo := DirInfo{file.Name(), dirSize}

				dirs = append(dirs, dirInfo)
				totalSize += dirSize
			}
		} else {
			rootDirSize += info.Size()
			totalSize += info.Size()
		}
	}

	dirs = append(dirs, DirInfo{"Root Directory", rootDirSize})
	return dirs, totalSize
}

func main() {
	var appArgs ProgramArgs

	p := arg.MustParse(&appArgs)
	dirName := appArgs.DirName

	if dirName == "" {
		dirName, _ = os.Getwd()
	}

	if appArgs.SortBy != "size" && appArgs.SortBy != "name" {
		p.Fail("--sort-by value must be either 'size' or 'name'")
	}

	if appArgs.SortOrder != "ascending" && appArgs.SortOrder != "descending" {
		p.Fail("--sort-order value must be either 'ascending' or 'descending'")
	}

	if appArgs.ListRootFilesOnly {
		dirs, totalSize := getListOfFiles(dirName, appArgs.ListAll)
		outputTable(sortDirList(dirs, appArgs.SortBy, appArgs.SortOrder), totalSize)
	} else {
		dirs, totalSize := getListOfDirs(dirName, appArgs.ListAll)
		outputTable(sortDirList(dirs, appArgs.SortBy, appArgs.SortOrder), totalSize)
	}
}
