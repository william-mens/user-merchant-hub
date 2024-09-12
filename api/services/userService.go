package services

import (
	"fmt"

	"bliss.com/tfcatalogue/entities"
	config "bliss.com/tfcatalogue/internal/database"
	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/google/uuid"
)

func getAll() {
	fmt.Println("hello world")
}

func Save(request *helpers.SetupUsers) (*entities.User, error) {
	userId := uuid.New().String()

	user := entities.User{
		Id:                 userId,
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		Email:              request.Email,
		Status:             "inactive",
		RegistrationStatus: "inactive",
	}

	result := config.Database.Where(entities.User{Email: request.Email}).FirstOrCreate(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
