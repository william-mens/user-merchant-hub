package helpers

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type Users struct {
	FirstName   string
	LastName    string
	Email       string
	Credentials []webauthn.Credential
}

func (u Users) WebAuthnID() []byte {
	userId := uuid.New()

	byteId, err := userId.MarshalBinary()
	if err != nil {
		fmt.Println("an error occurred converting uuid to byte array", err)
	}
	fmt.Printf("UUID as []byte: %v\n", byteId)
	return byteId

}

func (u Users) WebAuthnName() string {
	// fmt.spt
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u Users) WebAuthnDisplayName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u Users) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func UserCredentials() (*Users, error) {
	users := Users{
		FirstName:   "Willy",
		LastName:    "Bliss",
		Email:       "willybliss@gmail.com",
		Credentials: []webauthn.Credential{},
	}

	webAuthId := users.WebAuthnID()
	fmt.Printf("WebAuthnID: %v\n", webAuthId)

	// Example: Get WebAuthnName
	fmt.Printf("WebAuthnName: %s\n", users.WebAuthnName())
	fmt.Printf("WebAuthnDisplayName: %s\n", users.WebAuthnDisplayName())
	fmt.Printf("WebAuthnCredentials: %v\n", users.WebAuthnCredentials())
	return &users, nil

}
