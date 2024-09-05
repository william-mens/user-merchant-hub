package entities

import (
	"encoding/json"
	"time"
)

type MerchantProduct struct {
	Id            string           `json:"id" gorm:"type:uuid;primaryKey"`
	ProductId     string           `json:"productId"`
	MerchantId    string           `json:"merchantId"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Status        string           `json:"status" gorm:"default:active"`
	Configuration *json.RawMessage `json:"configuration" gorm:"type:json"`
	CreatedAt     *time.Time       `json:"createdAt"`
	UpdatedAt     *time.Time       `json:"updatedAt"`
}
