package luhn

import (
    "strconv"
    "strings"
)

func Valid(id string) bool {
    id = strings.TrimSpace(id)

    if len(id) <= 1 {
        return false
    }

    counter := 0
    sum := 0

    for idx := len(id) - 1; idx >= 0; idx-- {
        char := string(id[idx])

        if char == " " {
            continue
        }

        val, err := strconv.Atoi(char)
        if err != nil {
            return false
        }

        counter++

        if counter%2 == 0 {
            val *= 2

            if val > 9 {
                val -= 9
            }
        }

        sum += val
    }

    return counter > 1 && sum%10 == 0
}