package main

import "sort"

func sortDirList(dirs []DirInfo, sortType string, descending bool) []DirInfo {
	switch sortType {
	case "size":
		return sortBySize(dirs)
	case "name":
		return sortByName(dirs)
	default:
		return sortBySize(dirs)
	}
}

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
