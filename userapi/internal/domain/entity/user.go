package entity

type User struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Type    int       `json:"type"`
	Active  bool      `json:"active"`
	Address []Address `json:"addresses,omitempty"`
}

type Address struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Number       string `json:"number,omitempty"`
	Street       string `json:"street,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
	Code         string `json:"code,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
}
