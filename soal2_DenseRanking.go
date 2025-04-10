package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DenseRanking(scores []int, gits []int) []int {
	// Hilangkan duplikat skor untuk peringkat
	unique := []int{}
	seen := map[int]bool{}
	for _, score := range scores {
		if !seen[score] {
			unique = append(unique, score)
			seen[score] = true
		}
	}

	// Untuk setiap skor GITS, cari ranking-nya
	result := []int{}
	for _, gitsScore := range gits {
		rank := 1
		for _, s := range unique {
			if gitsScore < s {
				rank++
			} else {
				break
			}
		}
		result = append(result, rank)
	}

	return result
}

// Fungsi bantu konversi string ke []int
func parseInputToIntArray(input string) []int {
	parts := strings.Fields(input)
	arr := []int{}
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			arr = append(arr, num)
		}
	}
	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah pemain: ")
	reader.ReadString('\n') // Tidak dipakai, karena panjang skor bisa dihitung otomatis

	fmt.Print("Masukkan skor pemain (pisahkan dengan spasi): ")
	scoreInput, _ := reader.ReadString('\n')
	scores := parseInputToIntArray(scoreInput)

	fmt.Print("Masukkan jumlah permainan GITS: ")
	reader.ReadString('\n') // Tidak dipakai juga

	fmt.Print("Masukkan skor GITS (pisahkan dengan spasi): ")
	gitsInput, _ := reader.ReadString('\n')
	gits := parseInputToIntArray(gitsInput)

	result := DenseRanking(scores, gits)
	fmt.Print("Output: ")
	for _, rank := range result {
		fmt.Printf("%d ", rank)
	}
	fmt.Println()
}
