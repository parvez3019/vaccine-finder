package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var authHeader = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiSW5kaXZpZHVhbCIsInVzZXJJZCI6InRvVUNwUzlYcGo1dVlZUFlOb0tOMEE9PSIsImVtYWlsIjoiZDJnRGEzaFArZitSaDlVRTNQdW8yR2xtWEtZTkVySkZ0aXJBbGZXSkw4WVZyODZnSWhMRU55TDNQSFZodnFudCIsIm5iZiI6MTYzNDEyOTIzOCwiZXhwIjoxNjM0MTM1MjM4LCJpYXQiOjE2MzQxMjkyMzh9.B-3mEZM6yaeME5wbG2zwjlBFS4vPtc_BCNBfnpyYOK0"
var path = "/appointment/slots?countryCode=ind&missionCode=deu&centerCode=DEL&loginUser=pha3019%40gmail.com&visaCategoryCode=Blue%20Card%20with%20dependents&languageCode=en-US&applicantsCount=1&days=90&fromDate=14%2F10%2F2021&slotType=2&toDate=12%2F01%2F2022"

func main() {
	for {
		findSlot()
		time.Sleep(1 * time.Minute)
	}
}

func findSlot() {
	resp, err := fetchResponse()
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		notify("Login Required", "Re-login")
		return
	}

	defer resp.Body.Close()
	var response []Slot
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response ", err.Error())
		return
	}
	if len(response) > 1 {
		fmt.Printf("Slot Found %+v \n", response)
		notify("Slot Found", "Hurry up")
	} else if response[0].Error == nil {
		fmt.Printf("Slot Found %+v \n", response)
		notify("Slot Found", "Hurry up")
	} else {
		fmt.Printf("No slots %+v \n", response[0].Error)
	}
}

func notify(title, subtitle string) {
	cmd := Notification{
		Title:    title,
		Subtitle: subtitle,
		Message:  "",
	}
	cmd.Push()
}

func fetchResponse() (*http.Response, error) {
	url := "https://lift-api.vfsglobal.com" + path
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Accept-Language", "hi_IN")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	req.Header.Set("authorization", authHeader)
	req.Header.Set("route", "ind/en/deu")
	req.Header.Set("origin", "https://visa.vfsglobal.com")
	req.Header.Set("referer", "https://visa.vfsglobal.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error while fetching")
	}
	return resp, err
}

type Slot struct {
	Mission      interface{} `json:"mission"`
	Center       interface{} `json:"center"`
	Visacategory interface{} `json:"visacategory"`
	Date         interface{} `json:"date"`
	IsWeekend    bool        `json:"isWeekend"`
	Counters     interface{} `json:"counters"`
	Error        *ErrorModel `json:"error"`
}

type ErrorModel struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
