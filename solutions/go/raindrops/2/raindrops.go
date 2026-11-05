package raindrops
import (
	"strconv"
    "fmt"
)

func Convert(number int) string {

	result := ""
    /*
    // This approach is not correct because is go maps do not respect an order, so if I expect a result "PlingPlong" I could get "PlongPling"
    
    sounds := map[int]string{
        3: "Pling",
        5: "Plang",
        7: "Plong",	
	}

    for key, sound := range sounds {
    	if number%key == 0 {
        fmt.Println(strconv.Itoa(key) + " : " + sound)
    	    result = result + sound
    	}
	}*/
    
    rules := []struct {
        factor int
        sound  string
    }{
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    for _, rule := range rules {
        if number%rule.factor == 0 {
            result += rule.sound
        }
    }

	if result == "" {
	    result = strconv.Itoa(number)
	}

	fmt.Println(result)
    return result
}
