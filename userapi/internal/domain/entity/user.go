package entity

type Page struct {
	Paging  Paging `json:"paging"`
	Results []User `json:"results"`
}

type Paging struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

type User struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Active  bool      `json:"active"`
	Types   []Type    `json:"types,omitempty"`
	Address []Address `json:"addresses,omitempty"`
}

type Type struct {
	Type   string `json:"type"`
	UserId string `json:"-"`
}

type Address struct {
	Id           string `json:"-"`
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
