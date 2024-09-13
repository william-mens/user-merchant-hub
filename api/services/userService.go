package services

import (
	"encoding/json"
	"fmt"

	"bliss.com/tfcatalogue/entities"
	config "bliss.com/tfcatalogue/internal/database"
	"bliss.com/tfcatalogue/internal/helpers"
)

{
	"challenge": "K5IE0RIfTu_gRlfRaAWzNiFuY80DHAP8615sbOiU1ZI",
	"rpId": "go-webauthn.local",
	"user_id": "nkXX9qdQSrevsWFW/4li/A==",
	"expires": "0001-01-01T00:00:00Z",
	"userVerification": ""
}

type SessionData struct {
	Challenge string `json:"challenge"`
	RpId      string `json:"rpId"`
	UserId  string  `json:"user_id"`
	Expires  string    `json:"expires"`
	UserVerification string  `json:"userVerification"`
}

func getAll() {
	fmt.Println("hello world")
}

func Save(request *helpers.SetupUsers) (*entities.User, error) {

	user := entities.User{
		Id:                 request.Id,
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		Email:              request.Email,
		Status:             "inactive",
		RegistrationStatus: "inactive",
		Session:            request.Session,
	}

	result := config.Database.Where(entities.User{Email: request.Email}).FirstOrCreate(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func lookupUser(userId string) (*entities.User, error) {
	var user entities.User
	result := config.Database.Where(entities.User{Id: userId}).First(&user)
	if result.Error != nil {
		return nil, result.Error

	}
	var sessionJson SessionData
	err := json.Unmarshal([]byte(user.Session))

	return &user, nil
}
