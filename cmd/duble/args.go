package main

type ProgramArgs struct { // TODO: Make ProgramArgs
	DirName           string `arg:"positional"`
	ListRootFilesOnly bool   `arg:"-r, --root-only" default:"false" help:"List only the files in root directory/the passed directory name."`
	ListAll           bool   `arg:"-a, --list-all" default:"false" help:"Additionally list hidden files and directories."`
}
