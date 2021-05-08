package main

type Sessions struct {
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

type Centres struct {
	Name     string     `json:"name"`
	Address  string     `json:"address"`
	Sessions []Sessions `json:"sessions"`
}

type Availability struct {
	Centres []Centres `json:"centers"`
}
