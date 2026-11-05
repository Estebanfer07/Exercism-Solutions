package hamming

import (
    "fmt"
    "errors"
)

func Distance(a, b string) (int, error) {

    if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}
    
    distance := 0
    for idx, char := range a {
        fmt.Println(string(char))
	    if string(char) != string(b[idx]){
            distance++
        }
	}
    return distance, nil
}
