package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rumus(n int) string {
	var result []string
	current := 1

	for i := 0; i < n; i++ {
		result = append(result, fmt.Sprintf("%d", current))
		current += i + 1
	}

	return strings.Join(result, "-")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan angka: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input) //parse ke integer dari string
	if err != nil || n < 1 {
		fmt.Println("Input ngga valid masukkan angka bulat positif")
		return
	}

	output := rumus(n)
	fmt.Println("Output:", output)
}
