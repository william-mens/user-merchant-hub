package helpers

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type Users struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Credentials []webauthn.Credential
}

func (u Users) WebAuthnID() []byte {
	//converting a string to byeArray
	var authId []byte
	if len(u.Id) > 0 {
		//convert the string to byte array
		authId = []byte(u.Id)
	}
	authId, err := uuid.New().MarshalBinary()
	if err != nil {
		fmt.Println("an error occurred coverting uuid to byte", err)
		return []byte{}
	}
	return authId
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

func UserCredentials(request *SetupUsers) (webauthn.User, error) {
	users := Users{
		Id:          request.Id,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
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
