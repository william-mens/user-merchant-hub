package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"bliss.com/tfcatalogue/api/services"
	"bliss.com/tfcatalogue/entities"
	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
)

type registerReponse struct {
	Options *protocol.CredentialCreation
	Session *webauthn.SessionData
	User    *entities.User
}

func BeginRegistration(c *fiber.Ctx) error {
	createUser := new(helpers.SetupUsers)
	if err := c.BodyParser(createUser); err != nil {
		return helpers.BadRequestResponse(c, []string{"Invalid request payload"})
	}

	wconfig := &webauthn.Config{
		RPDisplayName: "Go Webauthn",                               // Display Name for your site
		RPID:          "go-webauthn.local",                         // Generally the FQDN for your site
		RPOrigins:     []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println("error with the wconfig", err)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	authAuth, authError := helpers.UserCredentials(createUser)
	if authError != nil {
		fmt.Println("error with auth Error occurred", authError)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	options, session, err := webAuthn.BeginRegistration(authAuth)
	if err != nil {
		fmt.Println("begin registration error occurred", err)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}
	userId := base64.StdEncoding.EncodeToString(session.UserID)

	createUser.Id = userId

	sessionJson, err := json.Marshal(session)
	if err != nil {
		fmt.Println("failed to convert data to json", err)
	}
	createUser.Session = string(sessionJson)

	user, dbErr := services.Save(createUser)

	if dbErr != nil {
		fmt.Println("an error occurred saving record in db", dbErr)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	response := registerReponse{
		options,
		session,
		user,
	}
	return helpers.SuccessResponse(c, "Success", response)

}
