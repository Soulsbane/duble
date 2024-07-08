package main

import "sort"

// Create a function to sort a slice of DirInfo by name.
func sortByName(dirs []DirInfo) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name < dirs[j].Name
	})

	return dirs
}

// Create a function to sort a slice of DirInfo by size.
func sortBySize(dirs []DirInfo) []DirInfo {
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size > dirs[j].Size
	})

	return dirs
}
