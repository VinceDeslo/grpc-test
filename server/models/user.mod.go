package models

type User struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Age         int32  `json:"Age"`
	Permissions string `json:"Permissions"`
}
