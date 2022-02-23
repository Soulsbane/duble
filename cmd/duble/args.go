package main

type args struct {
	DirName           string `arg:"positional"`
	ListRootFilesOnly bool   `arg:"-r, --root-only" default:"false" help:"List only the files in root directory/the passed directory name."`
}
