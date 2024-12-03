package main

import (
	"errors"
	"fmt"
	"sort"
)

// FilterAndSort filters numbers greater than or equal to the threshold and sorts them.
func FilterAndSort(nums []int, threshold int) []int {
	var filtered []int

	// Filter numbers greater than or equal to the threshold
	for _, num := range nums {
		if num >= threshold {
			filtered = append(filtered, num)
		}
	}

	// Sort the filtered numbers in ascending order
	sort.Ints(filtered)

	return filtered
}

// FindMostFrequent finds the most frequent word in a slice of strings.
// Returns an error if the slice is empty.
func FindMostFrequent(words []string) (string, error) {
	// Check if the input slice is empty
	if len(words) == 0 {
		return "", errors.New("input slice is empty")
	}

	// Count word frequencies
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Find the most frequent word
	var mostFrequentWord string
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			mostFrequentWord = word
			maxCount = count
		}
	}

	return mostFrequentWord, nil
}

func main() {
	// Test FilterAndSort
	fmt.Println("Testing Filter And Sort:")
	nums := []int{3, 10, 1, 7, 8, 2}
	threshold := 5
	fmt.Printf("Input: %v, Threshold: %d, Output: %v\n", nums, threshold, FilterAndSort(nums, threshold))

	// Test FindMostFrequent
	fmt.Println("\nTesting Find Most Frequent:")
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result, err := FindMostFrequent(words)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", words, result)
	}

	// Edge case: empty slice
	emptyWords := []string{}
	result, err = FindMostFrequent(emptyWords)
	if err != nil {
		fmt.Printf("Error: %v (empty input case handled)\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", emptyWords, result)
	}
}
