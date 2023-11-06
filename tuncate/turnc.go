package main

import (
	"fmt"
	"strings"
)

func main() {
	var number float32
	var count int
	fmt.Println("write a float number : ")
	fmt.Scan(&number)
	fmt.Println("count ")
	fmt.Scan(&count)
	Turncate(float64(number), count)
}
func Turncate(f float64, count int) {
	stringValue := fmt.Sprintf("%v", f)
	newStringValue := strings.Split(stringValue, "")
	dotBefore := []string{}
	i := 0
	for newStringValue[i] != "." {
		dotBefore = append(dotBefore, newStringValue[i])
		i++
	}

	var dotAfter []string = newStringValue[i:]
	var dotAfterTurn []string = dotAfter[:count+1]
	dotBefore = append(dotBefore, dotAfterTurn...)
	var stringTurn = strings.Join(dotBefore, "")
	// floatTurn, _ := strconv.ParseFloat(stringTurn, 32)
	fmt.Println(stringTurn)
}

// func roundFloat(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }
