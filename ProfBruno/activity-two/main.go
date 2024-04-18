package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mergeSort(data []int) []int {
	if len(data) <= 1 { // Here we are checking with the lenght of the slice is less or equal to one, if it's true, we have a single element, returning the data itself.
		return data
	}
	midNum := len(data) / 2          // This is the middle number of Merge Sort, here we are taking the lenght of "data" and dividing by two, in order to compose the middle number.
	rNum := mergeSort(data[midNum:]) // This is the right number of Merge Sorte, here we are selecting elements from the end of the slice.
	lNum := mergeSort(data[:midNum]) // This is the left number of Merge Sorte, here we are selecting elements from the beggining of the slice.
	return merge(rNum, lNum)
}

func merge(rNum, lNum []int) []int {
	result := make([]int, 0, len(rNum)+len(lNum))
	i, j := 0, 0
	for j < len(rNum) && i < len(lNum) {
		if lNum[i] < rNum[j] {
			result = append(result, lNum[i]) // This is adding lNum to result.
			i++
		} else {
			result = append(result, rNum[j]) // This is adding rNum to result.
			j++
		}
	}
	result = append(result, lNum[i:]...) // Appends all the elements fro lNum to the result.
	result = append(result, rNum[j:]...) // Appends all the elements fro rNum to the result.
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter with comma-separeted numbers: ")
	usrInpt, _ := reader.ReadString('\n')
	usrInpt = strings.TrimSpace(usrInpt)
	stringNumbers := strings.Split(usrInpt, ",")

	var data []int
	for _, strNum := range stringNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Error converting:", strNum)
			continue // Skip invalid input.
		}
		data = append(data, num)
	}
	sortedData := mergeSort(data)
	fmt.Println("The output: ", sortedData)
}
