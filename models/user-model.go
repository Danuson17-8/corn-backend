package models

type Translate struct {
	Th string `json:"th"`
	En string `json:"en"`
}

type User struct {
	Account string `json:"account"` // email
	Name    string `json:"name"`
	Role    string `json:"role"`
	Address string `json:"address"`
	Contact struct {
		Email       string `json:"email"`
		MobilePhone string `json:"mobile_phone"`
	} `json:"contact"`
}
