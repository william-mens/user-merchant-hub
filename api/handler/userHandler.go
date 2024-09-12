package handler

import (
	"fmt"

	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
)

type registerReponse struct {
	Options *protocol.CredentialCreation
	Session *webauthn.SessionData
}

func BeginRegistration(c *fiber.Ctx) error {

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

	authAuth, authError := helpers.UserCredentials()
	if authError != nil {
		fmt.Println("error with auth Error occurred", authError)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	options, session, err := webAuthn.BeginRegistration(authAuth)
	if err != nil {
		fmt.Println("begin registration error occurred", err)
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	response := registerReponse{
		options,
		session,
	}
	return helpers.SuccessResponse(c, "Success", response)

}
