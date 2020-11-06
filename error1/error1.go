package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	value, err := sqrt(9)
	if err != nil {
		fmt.Println("error")
	}else {
		fmt.Println(value)
	}

	value1, err := sqrt(-9)
	if err != nil {
		fmt.Println("error")
	}else {
		fmt.Println(value1)
	}
	contentType, err := contentTest("https://www.google.com")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(contentType)
	}
}

func sqrt(x float64) (float64, error) { //1. return 2 values , error object
	if x < 0 {
		return 0.0, fmt.Errorf("less than 0 (%f)", x)
	}
	return math.Sqrt(x), nil
}

func contentTest(url string)  (string, error){
	response, err := http.Get(url)
	if err != nil { // check if error
		return "", err
	}
	// no error then close the body and get required body
	defer response.Body.Close() // close the body

	contentLength := response.Header.Get("Content-Type")
	if contentLength == "" {
		return "", fmt.Errorf("cant find content Type")
	}else {
		return contentLength, nil
	}
}
