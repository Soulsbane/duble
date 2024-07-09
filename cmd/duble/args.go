package main

type ProgramArgs struct {
	DirName           string `arg:"positional"`
	ListRootFilesOnly bool   `arg:"-r, --root-only" default:"false" help:"List only the files in root directory/the passed directory name."`
	ListAll           bool   `arg:"-a, --list-all" default:"false" help:"Additionally list hidden files and directories."`
	SortBy            string `arg:"-s, --sort-by" default:"size" help:"Sort by size or name."`
	Descending        bool   `arg:"-d, --descending" default:"false" help:"Sort in descending order."`
}
