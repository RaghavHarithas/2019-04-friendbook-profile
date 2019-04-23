package models

type Profile struct {
	UserName     string `json: "username,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	Gender       string `json:"gender,omitempty"`
	DateOfBirth  string `json:"dateofbirth,omitempty"`
	MobileNumber string `json:"mobileNumber, omitempty"`
	City         string `json:"city,omitempty"`
}
