package services

import (
	"fmt"
	"strconv"

	"bliss.com/tfcatalogue/entities"
	config "bliss.com/tfcatalogue/internal/database"
	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/google/uuid"
)

type Pagination struct {
	CurrentPage  int `json:"currentPage"`
	PreviousPage int `json:"previousPage"`
	NextPage     int `json:"nextPage"`
	TotalRecords int `json:"totalRecords"`
	FromPage     int `json:"fromPage"`
	ToPage       int `json:"toPage"`
	TotalPages   int `json:"totalPages"`
	PerPage      int `json:"perPage"`
}

type DBResponse[T any] struct {
	Result   []T `json:"result"`
	Metadata struct {
		Pages helpers.Pagination `json:"pages"`
	} `json:"metadata"`
}

func GetAll(requests map[string]string) (DBResponse[entities.GetMerchants], error) {
	var merchants []entities.GetMerchants
	var count int64
	var res DBResponse[entities.GetMerchants]

	query := config.Database.Model(&entities.Merchant{})
	limit, err := strconv.Atoi(requests["limit"])
	page, pageErr := strconv.Atoi(requests["page"])
	if err != nil || limit == 0 {
		limit = 50
	}

	if pageErr != nil || page == 0 {
		page = 1

	}

	if len(requests) > 0 {
		for key, value := range requests {
			switch key {
			case "merchantId":
				query = query.Where("id = ?", value)
			case "alias":
				query = query.Where("alias = ?", value)
			case "companyName":
				query = query.Where("company_name = ?", value)
			case "code":
				query = query.Where("code = ?", value)
			}

		}
	}

	if err := query.Count(&count).Error; err != nil {
		return DBResponse[entities.GetMerchants]{}, fmt.Errorf("an error occurred querying total records: %w", err)
	}

	result := query.Scopes(
		helpers.Paginate(page, limit),
	).Find(&merchants)
	if result.Error != nil {
		return DBResponse[entities.GetMerchants]{}, fmt.Errorf("failed to retrieve merchants: %w", result.Error)
	}

	paginationMetadata := helpers.GetPaginationMetadata(page, int(count), limit)
	fmt.Println("final testPage", paginationMetadata)

	res.Result = merchants
	res.Metadata.Pages = paginationMetadata
	return res, nil
}

func Store(request *helpers.SetupMerchant) (*entities.Merchant, error) {

	newMerchantId := uuid.New()
	merchants := entities.Merchant{
		Id:          newMerchantId.String(),
		CompanyName: request.CompanyName,
		Code:        request.Code,
		TradeName:   stringOrNil(&request.TradeName),
		Alias:       request.Alias,
		Country:     request.Country,
		Status:      request.Status,
	}

	query := config.Database
	result := query.Create(&merchants)
	if result.Error != nil {
		return nil, result.Error
	}

	return &merchants, nil

}

func Update(request *helpers.SetupMerchant) (*entities.Merchant, error) {

	merchant := entities.Merchant{
		Id:          request.Id,
		CompanyName: request.CompanyName,
		Code:        request.Code,
		TradeName:   stringOrNil(&request.TradeName),
		Alias:       request.Alias,
		Country:     request.Country,
		Status:      request.Status,
	}
	result := config.Database.Model(&merchant).Updates(merchant)
	fmt.Println("updates full response")
	if result.Error != nil {
		return nil, result.Error
	}
	return &merchant, nil

}

func Destroy(merchantId string) (*entities.Merchant, error) {
	var merchant entities.Merchant

	// Retrieve the record first
	result := config.Database.Where("id = ?", merchantId).First(&merchant)
	if result.Error != nil {
		fmt.Println("first run the query", result.Error)
		return nil, result.Error
	}

	// Delete the record
	result = config.Database.Delete(&merchant)
	if result.Error != nil {
		fmt.Println("second run the query", result)
		return nil, result.Error
	}

	return &merchant, nil
}

func stringOrNil(s *string) *string {
	if *s == "" {
		return nil
	}
	return s
}
