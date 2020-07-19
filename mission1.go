package main

import (
	"fmt"
)

func main() {
	var numberOfTests int
	fmt.Scanf("%d\n", &numberOfTests)
	result := make([]int, numberOfTests)
	handleIO(numberOfTests,result,0)
	printResults(result, numberOfTests, 0)
}

func handleIO(numberOfTests int, result []int, outputIterator int) {
	if numberOfTests == 0 {
		return
	}

	var inputCount int
	fmt.Scanf("%d\n", &inputCount)
	inputNumbers := make([]int, inputCount)
	readNum(inputNumbers, 0, inputCount)

	var sum int = 0 
	var arrayLength int = inputCount 
	inputCount = arrayLength
	computeResults(inputNumbers, inputCount-1, &sum)
	result[outputIterator] = sum

	handleIO(numberOfTests-1, result, outputIterator+1)

}

func readNum(inputNumbers []int, start int, numberOfTests int){
	if numberOfTests == 0 {
		return
	}
	fmt.Scan(&inputNumbers[start])
	readNum(inputNumbers, start+1, numberOfTests-1)
}

func computeResults(inputNumbers []int, inputIterator int , sum *int) {
	if inputIterator == -1 {
		return
	}
	if inputNumbers[inputIterator] > 0 {
		*sum += inputNumbers[inputIterator]*inputNumbers[inputIterator]
	} 
	computeResults(inputNumbers, inputIterator-1, sum)
}

func printResults(result []int, numberOfTests int, outputIterator int) {
	if numberOfTests == 0 {
		return
	}
	fmt.Printf("%d\n",result[outputIterator])
	printResults(result, numberOfTests-1, outputIterator+1)
}