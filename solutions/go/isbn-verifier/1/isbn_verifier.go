package isbnverifier

import (
    "strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {

    normalizedIsbn := strings.ReplaceAll(isbn, "-", "")
    
    if len(normalizedIsbn) != 10 {
        return false
    }
    
	sum := 0
    multiplier := 10
    for idx, rune := range normalizedIsbn {
        char := string(rune)
        value := 0

        if char == "X" && idx == len(normalizedIsbn) - 1{
            value = 10
        } else {
    	    val, err := strconv.Atoi(char)
            if err != nil {
                return false
            }
            value = val
        }
        
        sum += value * multiplier
        multiplier--
	}

    return sum % 11 == 0
}
