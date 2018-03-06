package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/gorilla/mux"
)

var c = cache.New(1*time.Minute, 1*time.Minute)

func main() {

	router := mux.NewRouter()
	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("REIN", "reintoken", 30*time.Second)

	fmt.Println("Running Listening")
	router.HandleFunc("/health", GetHealth).Methods("GET")
	router.HandleFunc("/handleFirst", GetFirst).Methods("GET")
	router.HandleFunc("/handleSecond", GetSecond).Methods("GET")
	router.HandleFunc("/handleThird", GetThird).Methods("GET")
	router.HandleFunc("/handleFourth", GetFourth).Methods("GET")
	router.HandleFunc("/cache", GetCacheData).Methods("GET")
	router.HandleFunc("/handleFifth", GetFifth).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func GetCacheData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In get cache data function")
	foo, found := c.Get("REIN")
	if !found {
		fmt.Println("Expired Token  : Need to call auth and set the token")
		GetReinAuth()
	}
	fmt.Printf("Found : %t , Value : %s \n", found, foo)
}

func GetReinAuth() string {
	fmt.Println("Calling Rein Auth call")
	c.Set("REIN", "reintoken", 30*time.Second)

	token, _ := c.Get("REIN")
	t := reflect.TypeOf(token)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%+v\n", t.Field(i))
	}

	return "reintoken"
	//return token
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
