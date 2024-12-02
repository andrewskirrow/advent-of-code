package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	safeCount := 0
	err := readList(func(vals []int) {
		if isReadingSafe(vals) {
			safeCount++
		}
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("There are %d safe reports\n", safeCount)
}

func isReadingSafe(vals []int) bool {
	if len(vals) < 2 {
		return false
	}

	var mod int = 1
	if vals[0] < vals[1] {
		mod = -1
	}

	for i := 0; i < len(vals)-1; i++ {
		diff := mod * (vals[i] - vals[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func readList(collector func(vals []int)) error {
	f, err := os.Open("d2")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		itemsAsInt := make([]int, len(items))
		for i, v := range items {
			val, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return fmt.Errorf("Invalid number %s: %w", v, err)
			}

			itemsAsInt[i] = int(val)
		}

		collector(itemsAsInt)
	}

	return nil
}

