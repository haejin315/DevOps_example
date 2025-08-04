package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	A   int `json:"A"`
	B   int `json:"B"`
	Sum int `json:"Sum"`
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			log.Println("[CORS] preflight OPTIONS received")
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("요청 도착:", r.Method, r.URL.Path)

	err := r.ParseForm()
	if err != nil {
		log.Println("ParseForm 실패:", err)
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	aStr := r.FormValue("a")
	bStr := r.FormValue("b")
	log.Println("파라미터 a =", aStr, "b =", bStr)

	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)

	if errA != nil || errB != nil {
		log.Println("정수 변환 실패:", errA, errB)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := Result{A: a, B: b, Sum: a + b}
	log.Println("결과 생성:", result)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("JSON 인코딩 실패:", err)
	}
}

func main() {
	log.Println("서버 시작 중... http://localhost:8080")
	http.HandleFunc("/api/add", withCORS(AddHandler))
	http.ListenAndServe(":8080", nil)
}
