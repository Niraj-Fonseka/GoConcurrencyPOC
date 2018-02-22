package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health", GetHealth).Methods("GET")
	router.HandleFunc("/handleFirst", GetFirst).Methods("GET")
	router.HandleFunc("/handleSecond", GetSecond).Methods("GET")
	router.HandleFunc("/handleThird", GetThird).Methods("GET")
	router.HandleFunc("/handleFourth", GetFourth).Methods("GET")
	router.HandleFunc("/handleFifth", GetFifth).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Health")
	fmt.Fprintf(w, "Health")
}

func GetFirst(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In First")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "First")
}

func GetSecond(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Second")
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Second")
}

func GetThird(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Third")
	time.Sleep(3 * time.Second)

	fmt.Fprintf(w, "Third")
}

func GetFourth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Fourth")
	time.Sleep(5 * time.Second)

	fmt.Fprintf(w, "Fourth")
}

func GetFifth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Fifth")

	time.Sleep(6 * time.Second)

	fmt.Fprintf(w, "Fifth")
}
