package main

import "sort"

func sortByName(dirs []DirInfo) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name < dirs[j].Name
	})

	return dirs
}

func sortBySize(dirs []DirInfo) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size > dirs[j].Size
	})

	return dirs
}
