package modal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	Telephone string `json:"telephone"`
	NIC       string `json:"nic"`
	Age       uint   `json:"age"`
}
