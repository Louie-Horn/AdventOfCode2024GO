package main

import (
	"bufio"
	"fmt"
	"slices"

	//"io"
	"log"
	"os"
	//"strings"
)

func parseInput(path string) ([]string, error) {
	file, err := os.Open(path)
	var lines []string
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func firstStar(lines []string) (int, error) {
	var LeftList, RightList = []int{}, []int{}
	for _, line := range lines {
		var left, right int
		_, err := fmt.Sscanf(line, "%d   %d", &left, &right)
		if err != nil {
			return 1, err
		}
		LeftList = append(LeftList, left)
		RightList = append(RightList, right)
	}
	slices.Sort(LeftList)
	slices.Sort(RightList)
	diff := 0

	for i := 0; i < len(LeftList); i++ {
		if RightList[i] > LeftList[i] {
			diff += RightList[i] - LeftList[i]
		} else {
			diff += LeftList[i] - RightList[i]
		}
	}
	return diff, nil
}

func secondStar(lines []string) (int, error) {
	var LeftList, RightList = []int{}, []int{}
	for _, line := range lines {
		var left, right int
		_, err := fmt.Sscanf(line, "%d   %d", &left, &right)
		if err != nil {
			return 1, err
		}
		LeftList = append(LeftList, left)
		RightList = append(RightList, right)
	}
	similarity := 0
	for _, left := range LeftList {
		app_count := 0
		for _, right := range RightList {
			if left == right {
				app_count++
			}
		}
		similarity += app_count * left
	}
	return similarity, nil
}
func main() {
	lines, err := parseInput("../input-data/01.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	diff, err := firstStar(lines)
	if err != nil {
		log.Fatalf("firstStar: %s", err)
	}
	fmt.Printf("First Star: %d\n", diff)

	similarity, err := secondStar(lines)
	if err != nil {
		log.Fatalf("secondStar: %s", err)
	}
	fmt.Printf("Second Star: %d\n", similarity)
}
