package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
}

func reorganizeString(s string) string {
	n := len(s)
	if n == 0 || n == 1 {
		return s
	}
	dp := make(map[byte]int)

	max := 0
	maxChar := byte('0')
	for i := 0; i < n; i++ {
		dp[s[i]]++
		count, _ := dp[s[i]]
		if count > max {
			max = count
			maxChar = s[i]
		}
	}
	if max > (n+1)/2 {
		return ""
	}
	res := make([]byte, n)
	dp[maxChar] = 0
	i := 0
	for ; max > 0; i = i + 2 {
		res[i] = maxChar
		max--
	}
	for ch, count := range dp {
		for count > 0 {
			if i >= n {
				i = 1
			}
			res[i] = ch
			i = i + 2
			count--
		}
	}
	return string(res)
}
