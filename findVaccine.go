package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fetchVaccineDetails(fetchNext20Dates [20]string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	ids := pinCodes
	if DistrictBasedSearch {
		ids = districtIds
	}
	for _, id := range ids {
		for _, date := range fetchNext20Dates {
			availability := fetchVaccineRequest(id, date)
			if availability == nil {
				return
			}
			notify(availability, id, date)
		}
		time.Sleep(1 * time.Minute)
		fmt.Println("----------------------")
	}
}

func notify(availability *Availability, id string, date string) {
	found := false
	for _, centre := range availability.Centres {
		for _, session := range centre.Sessions {
			if session.MinAgeLimit == Age && session.AvailableCapacity > 0 {
				found = true
				cmd := Notification{
					Title:    "Vaccine Slot Available",
					Subtitle: fmt.Sprintf("%s - %s", session.Date, centre.Address),
					Message:  fmt.Sprintf("Count - %d", session.AvailableCapacity),
				}
				cmd.Push()
				fmt.Printf("Slot found : %+v - %+v", centre, session)
			}
		}
	}
	if !found {
		fmt.Printf("No vaccines found for date %s at pincode/district %s\n", date, id)
	}
}

func getDates() [20]string {
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, 2)
	//10-05-2021
	dates := [20]string{}
	for i := 0; i < 20; i++ {
		dates[i] = currentTime.Format("02-01-2006")
		currentTime = currentTime.AddDate(0, 0, 1)
	}
	return dates
}

func fetchVaccineRequest(id string, date string) *Availability {
	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=%s&date=%s", id, date)
	if DistrictBasedSearch {
		url = fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?district_id=%s&date=%s", id, date)
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Accept-Language", "hi_IN")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error while fetching")
	}

	defer resp.Body.Close()
	var availability Availability
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&availability)
	if err != nil {
		fmt.Println("Rate Limiter Exceeded")
		return nil
	}

	return &availability
}
