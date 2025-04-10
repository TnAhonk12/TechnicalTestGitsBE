package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsBalancedBracket(s string) string {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan deretan bracket (contoh: { [ ( ) ] }):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Hapus spasi jika ada
	cleaned := strings.ReplaceAll(input, " ", "")

	// Cek keseimbangan
	result := IsBalancedBracket(cleaned)
	fmt.Println("Output:", result)
}
