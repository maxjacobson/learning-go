package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	palindromes := []int{}

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			str := strconv.FormatInt(int64(product), 10)
			if isPalindrome(str) {
				palindromes = append(palindromes, product)
			}
		}
	}

	sort.Ints(palindromes)
	fmt.Println(palindromes[len(palindromes)-1])
}

func isPalindrome(str string) bool {
	chars := strings.Split(str, "")
	for i := 0; i < len(chars); i++ {
		if chars[i] != chars[len(chars)-i-1] {
			return false
		}
	}
	return true
}
