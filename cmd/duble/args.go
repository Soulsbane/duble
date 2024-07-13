package main

type ProgramArgs struct {
	DirName           string `arg:"positional"`
	ListRootFilesOnly bool   `arg:"-r, --root-only" default:"false" help:"List only the files in root directory/the passed directory name."`
	ListAll           bool   `arg:"-a, --list-all" default:"false" help:"Additionally list hidden files and directories."`
	SortBy            string `arg:"-s, --sort-by" default:"size" help:"Sort by size or name."`
	SortOrder         string `arg:"-o, --sort-order" default:"descending" help:"Sort in descending or ascending order."`
}

func (args ProgramArgs) Description() string {
	return "Duble shows the size of each directory below user specified directory or current working directory if none specified."
}

func (args ProgramArgs) Epilogue() string {
	return "For Duble project updates and bug reports, visit https://github.com/Soulsbane/duble"
}
