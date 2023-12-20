package domain

//user represents a user in the system

type User struct {
	ID int `gorm:"primarykey"`
	Name string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Username string `json:"username"`
    Password string `json:"password"`
	Phone string `gorm:"unique" json:"phone"`
	Permission bool `gorm:"default:true" json:"permission"`
}
// Address represnts the address of a user
type Address struct {
	ID uint `json:"id " gorm:"unique;not null"` 
}
