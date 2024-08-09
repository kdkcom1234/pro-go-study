package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		// i = i // 잠재적인 오류 발생 가능(go vet으로 확인)
		PrintNumber(i)
	}
}

// revivie:disable:exported
func PrintHello() {
	// 패캐지.함수명(매개변수)
	fmt.Println("Hello, Go")
}

// revive:enable:exported

// PrintNumber 함수는 fmt.Println 함수를 사용해서 숫자를 출력한다.
func PrintNumber(number int) {
	fmt.Println(number)
}
