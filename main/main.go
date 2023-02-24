package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TestCase struct {
	SawmillCount int
	SawmillTests []SawMillTest
}

type SawMillTest struct {
	TreeTrunkCount int
	TreeTrunks     []int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	testCases := []TestCase{}
	inputState := "start"
	fmt.Println("Enter test cases: ")
	currentInput := TestCase{SawmillTests: []SawMillTest{}}
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input == "0" {
			break
		}
		switch inputState {
		case "start":
			sawMillCount, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Error: %v\n Enter a valid text: \n", err)
				continue
			}
			currentInput.SawmillCount = sawMillCount
			inputState = "reading_trunks"
			continue
		case "reading_trunks":
			splitted_inputs := strings.Split(input, " ")
			trunksCount, err := strconv.Atoi(splitted_inputs[0])
			if err != nil {
				fmt.Printf("Error: %v\n Enter a valid text: \n", err)
				continue
			}
			sawMilltest := SawMillTest{TreeTrunkCount: trunksCount, TreeTrunks: make([]int, trunksCount)}
			for i := 0; i < trunksCount; i++ {
				trunkSize, err := strconv.Atoi(splitted_inputs[i+1])
				if err != nil {
					fmt.Printf("Error: %v\n Enter a valid text: \n", err)

					continue
				}
				sawMilltest.TreeTrunks[i] = trunkSize
			}
			currentInput.SawmillTests = append(currentInput.SawmillTests, sawMilltest)
			if len(currentInput.SawmillTests) == currentInput.SawmillCount {
				inputState = "start"
				testCases = append(testCases, currentInput)
				currentInput = TestCase{SawmillTests: []SawMillTest{}}
			}
			continue
		}

	}

	for i, testCase := range testCases {
		evaluateTestCase(testCase, i)
	}
}

func evaluateTestCase(testCase TestCase, index int) {
	fmt.Printf("Case %v\n", index+1)
	maxIncome := 0
	allOrders := [][][]int{}
	for _, test := range testCase.SawmillTests {
		currentMaxIncome, currentOrders := calculateBestOrders(test)
		maxIncome += currentMaxIncome
		allOrders = append(allOrders, currentOrders)
	}
	fmt.Printf("Max profit: %v\n", maxIncome)
	for i, orders := range allOrders {
		for _, orderedTrunks := range orders {
			fmt.Print(orderedTrunks)
		}
		if i == len(allOrders)-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(",")
		}
	}
}

// Returns maxIncome, Orders that generate that income
func calculateBestOrders(test SawMillTest) (int, [][]int) {
	allPossiblePermutations := permutations(test.TreeTrunks)
	currentMaxIncome := 0
	currentOrders := [][]int{}
	for _, orderedTrunks := range allPossiblePermutations {
		currentIncome := calculateIncome(orderedTrunks)
		if currentIncome > currentMaxIncome {
			currentMaxIncome = currentIncome
			currentOrders = [][]int{orderedTrunks}
		} else if currentIncome == currentMaxIncome && !includes(currentOrders, orderedTrunks) {
			currentOrders = append(currentOrders, orderedTrunks)
		}
	}
	return currentMaxIncome, currentOrders
}

/////////////
/////////////
//// I decided not to continue writing the following function because it was going to take too much time for me to implement
/// And I was not completely sure it was going to work. So I will go with the brute force option
// Which is to generate all permutations of the array and evaluate income for each
/////////////
/////////////

/*
// Calculates all possible orders
// The idea is that the best is to constantly make sure to have a remainder of 1 so then we will have
// a sawn wood with length 2
// so for example: 4,3,6,9,3,4,3,3,1
calculateOrders(treeTrunks []int) [][]int {

	maxTrunkSize := 0
	// First, lets get a map where key will be trunkSize and value, amount
	trunksBySize := map[int]int{}
	for _, trunk := range treeTrunks {
		trunksBySize[trunk] += 1
		if trunk > maxTrunkSize {
			maxTrunkSize = trunk
		}
	}


	orderedTrunks := [][]int{}
	for trunkSize, amount := range trunksBySize {
		switch trunkSize % 3 {
		case 1:
			// should be first
		}
	}
}
*/

func calculateIncome(treeTrunks []int) int {
	income := 0
	currentRemainder := 0
	for _, trunk := range treeTrunks {
		// If trunk + remainder is enough to make a cut
		if trunk+currentRemainder >= 3 {
			firstCut := 3 - currentRemainder
			income += incomeFor(firstCut)
			income += ((trunk - firstCut) / 3) * incomeFor(3)
		}
		currentRemainder = (trunk + currentRemainder) % 3
	}
	return income
}

func incomeFor(sawnWoodLength int) int {
	switch sawnWoodLength {
	case 1:
		return -1
	case 2:
		return 3
	case 3:
		return 1
	}
	return 0
}

// From https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func includes(arrayOfArrays [][]int, array []int) bool {
	for _, arrayElem := range arrayOfArrays {
		match := true
		for i, elem := range arrayElem {
			if array[i] != elem {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
