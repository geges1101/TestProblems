package main

import "strings"

func simplifyPath(path string) string {
	res := make([]string, 0)
	split := strings.Split(path, "/")

	for _, dir := range split {
		if dir == "" || dir == "." {
			continue
		}
		if dir != ".." {
			res = append(res, dir)

		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}

	return "/" + strings.Join(res, "/")
}
