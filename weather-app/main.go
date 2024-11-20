package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//For this to work you need an weatherapi from https://www.weatherapi.com/
//insert apikey to "APIKEYHERE"
//insert your city at " city := "Your-city-here" "
//you need to change the apikey and city on the URL to placeholders "%s"
//http://api.weatherapi.com/v1/timezone.json?key=**apikey**&q=**your city** -> ..key=%s&q=%s&aqi=yes"
//insert weatherapi URL at getweather same at getTime
//To make it into an EXE "go build -o weatherapp.exe main.go"

type weatherResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TzID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type 'run' to print weather and time info:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)

	if input == "run" {
		city := "YOUR-CITY-HERE"
		apikey := "APIKEYHERE"

		getWeather(city, apikey)
		getTime(city, apikey)
	} else {
		fmt.Println("Invalid input.")
		return
	}

	fmt.Println("Press Enter to quit...")
	_, _ = reader.ReadString('\n')
	cmd := exec.Command("taskkill", "/F", "/IM", "cmd.exe")
	cmd.Run()
}

func getWeather(city string, apikey string) {
	// URL for the weather API goes here ->
	url := fmt.Sprintf("WEATHERAPI URL HERE", apikey, city)

	// sends request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	// reads the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var weatherData weatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	//weather details
	fmt.Printf("The temperature in %s, %s (%s) is %.2f°C\n", weatherData.Location.Name, weatherData.Location.Region, weatherData.Location.Country, weatherData.Current.TempC)
	fmt.Printf("Local time: %s\n", weatherData.Location.Localtime)
}

func getTime(city string, apikey string) {
	// URL for the weather API goes here too ->
	url := fmt.Sprintf("WEATHERAPI URL HERE TOO", apikey, city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching time data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var timeData weatherResponse
	err = json.Unmarshal(body, &timeData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("Current local time in %s is %s\n", timeData.Location.Name, timeData.Location.Localtime)
}
