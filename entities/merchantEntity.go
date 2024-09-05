package entities

import "time"

type Merchant struct {
	Id               string            `json:"id" gorm:"type:uuid;primaryKey"`
	CompanyName      string            `json:"companyName"`
	Code             string            `json:"code"`
	TradeName        *string           `json:"tradeName"`
	Alias            string            `json:"alias"`
	Country          string            `json:"country" gorm:"default:GH"`
	Status           string            `json:"status" gorm:"default:active"`
	CreatedAt        *time.Time        `json:"createdAt"`
	UpdatedAt        *time.Time        `json:"updatedAt"`
	MerchantProducts []MerchantProduct `gorm:"foreignkey:MerchantId"`
}

type GetMerchants struct {
	Id          string     `json:"id" gorm:"type:uuid;primaryKey"`
	CompanyName string     `json:"companyName"`
	Code        string     `json:"code"`
	TradeName   *string    `json:"tradeName"`
	Alias       string     `json:"alias"`
	Country     string     `json:"country" gorm:"default:GH"`
	Status      string     `json:"status" gorm:"default:active"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
