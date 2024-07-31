package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors []string
}

// 슬라이스 생성
// make(타입, 초기크기, 용량)
// 용량은 사이즈 변경 없이 슬라이스에 추가할 수 있음
// []*Rsvp: 구조체 포인터 슬라이스
var responses = make([]*Rsvp, 0, 10)

// 맵 생성
// map[키타입]값타입
var templates = make(map[string]*template.Template, 3)

func loadTemplate() {
	// TODO - 템플릿 로딩
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

// func formHandler(writer http.ResponseWriter, request *http.Request) {
// 	if request.Method == http.MethodGet {
// 		// 값은 중괄호를 사용해 생성
// 		templates["form"].Execute(writer, formData{Rsvp: &Rsvp{}, Errors: []string{}})
// 	}
// }

func formGetHandler(writer http.ResponseWriter, request *http.Request) {
	// 값은 중괄호를 사용해 생성
	templates["form"].Execute(writer, formData{Rsvp: &Rsvp{}, Errors: []string{}})
}

func formPostHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	responseData := Rsvp{
		Name:       request.Form["name"][0],
		Email:      request.Form["email"][0],
		Phone:      request.Form["phone"][0],
		WillAttend: request.Form["willattend"][0] == "true",
	}

	errors := []string{}
	if responseData.Name == "" {
		errors = append(errors, "Please enter your name")
	}
	if responseData.Email == "" {
		errors = append(errors, "Please enter your email address")
	}
	if responseData.Phone == "" {
		errors = append(errors, "Please enter your phone number")
	}

	if len(errors) > 0 {
		templates["form"].Execute(writer, formData{Rsvp: &responseData, Errors: errors})
	}

	responses = append(responses, &responseData)
	if responseData.WillAttend {
		templates["thanks"].Execute(writer, responseData.Name)
	} else {
		templates["sorry"].Execute(writer, responseData.Name)
	}
}

func main() {
	// fmt.Println("TODO: add some features")
	loadTemplate()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("GET /form", formGetHandler)
	http.HandleFunc("POST /form", formPostHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
