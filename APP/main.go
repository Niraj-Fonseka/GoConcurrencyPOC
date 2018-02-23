package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Main Function Running")

	go runAllRequestsConc()

	fmt.Scanln()
}

func runAllRequests() {
	fmt.Println("Running Functions sequentially \n\n")

	// getFour()
	// getHealth()
	// getOne()
	// getFive()
	// getTwo()
	// getThree()

}

func runAllRequestsConc() {
	fmt.Println("Running Functions Concurrently \n\n")
	//health := make(chan string)
	One := make(chan string)
	Three := make(chan string)
	Four := make(chan string)
	Five := make(chan string)
	Two := make(chan string)

	go getFour(Four)
	go getFive(Five)
	go getTwo(Two)
	go getOne(One)
	go getThree(Three)
	go processWorkers(One, Five, Three, Four, Two)
	//fiveOut, ok := <-Five

	// go getTwo()
	// go getThree()

}

func processWorkers(channels ...<-chan string) {
	fmt.Println("In the process workers function")
	fmt.Printf("Length of the channels : %d \n ", len(channels))
	// select {
	// case x, ok := <-channels[1]:
	// 	if ok {
	// 		fmt.Printf("Value %x was read.\n", x)
	// 	} else {
	// 		fmt.Println("Channel closed!")
	// 	}
	// default:
	// 	fmt.Println("No value ready, moving on.")
	// }

	for index := 0; index < len(channels); index++ {
		x, ok := <-channels[index]
		//fmt.Printf("In the for loop , Index : %d \n", index)
		if ok {
			fmt.Printf("Value is here : %s \n", x)
		}
	}

	// value5, ok := <-channels[1]
	// fmt.Printf("Value : %s , Done : %t \n", value, ok)

}

func getHealth(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/health", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- string(body)

	fmt.Printf("GetHealth : %s  \n", body)
	return string(body)
}

func getOne(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFirst", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetOne : %s  \n", body)
	ch <- string(body)

	return string(body)
}

func getTwo(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleSecond", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetTwo : %s  \n", body)
	ch <- string(body)

	return string(body)
}

func getThree(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleThird", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetThree : %s  \n", body)
	ch <- string(body)

	return string(body)
}

func getFour(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFourth", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetFour : %s  \n", body)
	ch <- string(body)

	return string(body)
}

func getFive(ch chan<- string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000/handleFifth", nil)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetFive : %s  \n", body)
	ch <- string(body)

	return string(body)
}
