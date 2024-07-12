package main

import "sort"

func sortDirList(dirs []DirInfo, sortType string, sortOrder string) []DirInfo {
	switch sortType {
	case "size":
		return sortBySize(dirs, sortOrder)
	case "name":
		return sortByName(dirs, sortOrder)
	default:
		return sortBySize(dirs, sortOrder)
	}
}

func sortByName(dirs []DirInfo, sortOrder string) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		if sortOrder == "descending" {
			return dirs[i].Name > dirs[j].Name
		} else {
			return dirs[i].Name < dirs[j].Name
		}
	})

	return dirs
}

func sortBySize(dirs []DirInfo, sortOrder string) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		if sortOrder == "descending" {
			return dirs[i].Size > dirs[j].Size
		} else {
			return dirs[i].Size < dirs[j].Size
		}
	})

	return dirs
}
