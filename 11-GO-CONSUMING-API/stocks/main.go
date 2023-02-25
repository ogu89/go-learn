package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func main() {
	// response, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=IBM&apikey=demo")

	key := "NXRTXG0LJKJEY1GB"
	url := "https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=AAPL&apikey="+key

	responseApple, errApple := http.Get(url)


	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	if errApple != nil {
		fmt.Print(errApple.Error())
		os.Exit(1)
	}

	// responseData, err := ioutil.ReadAll(response.Body)
	responseAppleData, errApple := ioutil.ReadAll(responseApple.Body)


	// if err != nil {
	// 	log.Fatal(err)
	// }

	if errApple != nil {
		log.Fatal((errApple))
	}


	// fmt.Println(string(responseData))
	fmt.Println(string(responseAppleData))
}