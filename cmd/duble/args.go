package main

type args struct {
	DirName     string `arg:"positional"`
	ListSubDirs bool   `arg:"-l, --list" default:"false" help:"List directories under the passed directory."`
}
