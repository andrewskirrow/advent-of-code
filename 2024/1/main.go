package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	value1List := make([]int, 1000)
	value2Map := make(map[int]int, 1000)

	err := readList(func(i1, i2 int) {
		value1List = append(value1List, i1)
		count2 := value2Map[i2] + 1
		value2Map[i2] = count2
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	similarity := 0
	for _, i1 := range value1List {
		similarity = similarity + (i1 * value2Map[i1])
	}

	fmt.Printf("Similarity: %d\n", similarity)
}

func calculateDifference() {
	var (
		list1 []int = make([]int, 0, 2000)
		list2 []int = make([]int, 0, 2000)
	)

	readList(func(i1, i2 int) {
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	})

	slices.Sort(list1)
	slices.Sort(list2)

	diff := 0
	for i, e1 := range list1 {
		e2 := list2[i]
		var d int
		if e1 > e2 {
			d = e1 - e2
		} else {
			d = e2 - e1
		}

		diff += d
	}

	fmt.Printf("difference is: ", diff)
}

func readList(collector func(i1, i2 int)) error {
	f, err := os.Open("d1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), "   ")
		i1, err := strconv.ParseInt(txt[0], 10, 32)
		if err != nil {
			fmt.Printf("%s is not a  valid int: %s", txt[0], i1)
		}
		i2, err := strconv.ParseInt(txt[1], 10, 32)
		if err != nil {
			fmt.Printf("%s is not a  valid int: %s", txt[1], i2)
		}

		collector(int(i1), int(i2))
	}

	return nil
}

