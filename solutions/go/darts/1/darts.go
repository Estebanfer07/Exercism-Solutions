package darts

import "math"

func Score(x, y float64) int {
    scores := []struct {
        minRadius float64
        maxRadius float64
        score int
    }{
        {minRadius: 0, maxRadius: 1, score: 10},
        {minRadius: 1, maxRadius: 5, score: 5},
        {minRadius: 5, maxRadius: 10, score: 1},
    }
    hipotenuse := math.Sqrt((math.Pow(x, 2) + math.Pow(y, 2)))

    score := 0
    for _, value := range scores {
        if hipotenuse >= value.minRadius && hipotenuse <= value.maxRadius {
            score = value.score
            break
        }
    }

    return score
}
