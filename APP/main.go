package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Main Function Running")

	runAllRequestsConc()

	fmt.Scanln()
}

func runAllRequests() {
	fmt.Println("Running Functions sequentially \n\n")

	getFour()
	getHealth()
	getOne()
	getFive()
	getTwo()
	getThree()

}

func runAllRequestsConc() {
	fmt.Println("Running Functions Concurrently \n\n")

	go getFour()
	go getHealth()
	go getOne()
	go getFive()
	go getTwo()
	go getThree()

}

func getHealth() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/health", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("GetHealth : %s  \n", body)
	return string(body)
}

func getOne() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFirst", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetOne : %s  \n", body)

	return string(body)
}

func getTwo() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleSecond", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetTwo : %s  \n", body)

	return string(body)
}

func getThree() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleThird", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetThree : %s  \n", body)

	return string(body)
}

func getFour() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFourth", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetFour : %s  \n", body)

	return string(body)
}

func getFive() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFifth", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetFive : %s  \n", body)

	return string(body)
}
