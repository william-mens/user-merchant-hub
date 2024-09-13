package helpers

type (
	Merchant struct {
		MerchantId  string `json:"merchantId" validate:"omitempty,max=36"`
		CompanyName string `json:"companyName" validate:"omitempty,max=100"`
		Limit       string `json:"limit" validate:"omitempty,max=100"`
		Tags        string `json:"tags" validate:"omitempty,max=100"`
		Alias       string `json:"alias" validate:"omitempty,max=100"`
	}

	SetupMerchant struct {
		Id                        string   `json:"id" validate:"omitempty,max=36"`
		CompanyName               string   `json:"companyName" validate:"max=100"`
		Tags                      []string `json:"tags" validate:"dive,min=3,max=10"`
		Alias                     string   `json:"alias" validate:"omitempty,max=100"`
		Code                      string   `json:"code" validate:"omitempty,max=100"`
		Country                   string   `json:"country" validate:"omitempty,max=100"`
		TradeName                 string   `json:"tradeName" validate:"max=100"`
		CompanyRegistrationNumber string   `json:"companyRegistrationNumber" validate:"omitempty,max=15"`
		TypeOfCompany             string   `json:"typeOfCompany" validate:"omitempty,max=100"`
		VatRegistrationNumber     string   `json:"vatRegistrationNumber" validate:"omitempty,max=100"`
		DateOfIncorporation       string   `json:"dateOfIncorporation" validate:"omitempty,max=100"`
		DateOfCommencement        string   `json:"dateOfCommencement" validate:"omitempty,max=100"`
		CompanyLogo               string   `json:"companyLogo" validate:"omitempty,max=100"`
		TaxIdentificationNumber   string   `json:"taxIdentificationNumber" validate:"omitempty,max=100"`
		Status                    string   `json:"status" validate:"omitempty,max=100"`
	}

	GetMerchantProducts struct {
		MerchantId        string `json:"merchantId" validate:"omitempty,min=3,max=36"`
		MerchantProductId string `json:"merchantProductId" validate:"omitempty,min=3,max=36"`
	}

	SetupUsers struct {
		Id        string `json:"id" validate:"omitempty,max=36"`
		FirstName string `json:"firstName" validate:"omitempty,min=3,max=36"`
		LastName  string `json:"lastName" validate:"omitempty,min=3,max=36"`
		Email     string `json:"email" validate:"omitempty,min=3,max=36"`
		Session   string `json:"session" validate:"omitempty,max=36"`
	}
)
