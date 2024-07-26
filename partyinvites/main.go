package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// 슬라이스 생성
// make(타입, 초기크기, 용량)
// 용량은 사이즈 변경 없이 슬라이스에 추가할 수 있음

// []*Rsvp: 구조체 포인터 슬라이스
var responses = make([]*Rsvp, 0, 10)

func main() {
	fmt.Println("TODO: add some features")
}
