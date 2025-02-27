package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
}

func sortArr(arr []string) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func sortMap(m map[string]int) []string {
	arr := []string{}
	s := strings.Builder{}
	for i, v := range m {
		for j := 0; j < v; j++ {
			s.WriteString(i)
		}
		arr = append(arr, s.String())
		s.Reset()
	}
	return arr
}

func goMap(s string) map[string]int {
	mapka := make(map[string]int, len(s))
	for _, v := range s {
		if unicode.IsLower(v) {
			mapka[string(v)]++
		}
	}
	return mapka
}
