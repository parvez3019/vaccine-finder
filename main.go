package main

import (
	"time"
)

const Age = 18
const DistrictBasedSearch = false

// Choose either pin code based search or district
// Due to rate limiter on api add only up to 5 pins/district ids
var pinCodes = []string{
	"110075",
	"110077",
	"110022", // RK Puram
	"110058", // Janak Puri
	"110018", // Tilak Nagar
}

var districtIds = []string{
	"1111", // dummy id
}

func main() {
	// For next 20 days
	fetchNext20Dates := getDates()

	for {
		go func(fetchNext20Dates [20]string, pinCodes []string) {
			fetchVaccineDetails(fetchNext20Dates)
		}(fetchNext20Dates, pinCodes)
		time.Sleep(6 * time.Minute)
	}

}
