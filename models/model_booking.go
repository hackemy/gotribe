package models

type Booking struct {
	Seats        int    `json:"seats"`
	Start        string `json:"start"`
	End          string `json:"end"`
	DisplayStart string `json:"displayStart"`
	DisplayEnd   string `json:"displayEnd"`
	State        string `json:"state"`
}
