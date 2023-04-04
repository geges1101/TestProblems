package main

func generateParenthesis(n int) []string {
	var res []string

	helper(n, 0, 0, "", &res)

	return res
}

func helper(n int, openCount int, closeCount int, currentStr string, result *[]string) {
	if openCount == n && closeCount == n {
		*result = append(*result, currentStr)
		return
	}

	if openCount < n {
		helper(n, openCount+1, closeCount, currentStr+"(", result)
	}

	if openCount > closeCount {
		helper(n, openCount, closeCount+1, currentStr+")", result)
	}
}
