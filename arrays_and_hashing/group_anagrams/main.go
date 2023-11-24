package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	if len(strs) == 1 {
		return [][]string{strs}
	}

	hashmap := map[string][]string{}

	for _, str := range strs {
		sortedString := sortString(str)
		hashmap[sortedString] = append(hashmap[sortedString], str)
	}

	res := [][]string{}
	for _, v := range hashmap {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	stringList := strings.Split(s, "")
	sort.Strings(stringList)
	return strings.Join(stringList, "")
}
