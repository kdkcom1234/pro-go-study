package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// fmt.Println("Hello, Operations")

	kayak := 275
	soccerBall := 19.50

	// total := float64(kayak) + soccerBall
	total := kayak + int(math.Round(soccerBall))
	fmt.Println(total)

	val1 := "10x0"
	int1, int1err := strconv.Atoi(val1)
	if int1err == nil {
		fmt.Println(int1)
	} else {
		// 에러일 때 0으로 초기화됨, 값이 아닌 에러로 체크
		fmt.Println("Cannot parse", int1, int1err)
	}

	val := 275
	base10string := strconv.Itoa(val)

	fmt.Println(base10string)
}
