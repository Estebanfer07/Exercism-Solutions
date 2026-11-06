package differenceofsquares

import (
    "math"
	"fmt"
)

func seq(n int) []int {
	if n <= 0 {
		return []int{}
	}
	// Recursively build the slice up to n-1, then append n
	return append(seq(n-1), n)
}

func SquareOfSum(n int) int {
	numbers := seq(n)
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	numbers := seq(n)
    sum := 0
    
    fmt.Println(numbers)
    for _, number := range numbers {
        sum += int(math.Pow(float64(number), 2))
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
