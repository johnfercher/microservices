package entity

type User struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Address  *Address `json:"address,omitempty"`
	Login    string   `json:"login"`
	Password string   `json:"password"`
}

type Address struct {
	Code string `json:"code"`
}
