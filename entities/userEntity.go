package entities

import "time"

type User struct {
	Id                 string     `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName          string     `json:"firstName"`
	LastName           string     `json:"lastName"`
	Email              string     `json:"email"`
	Status             string     `json:"status" gorm:"default:active"`
	RegistrationStatus string     `json:"registrationStatus" gorm:"default:inactive"`
	CreatedAt          *time.Time `json:"createdAt"`
	UpdatedAt          *time.Time `json:"updatedAt"`
}
