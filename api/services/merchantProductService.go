package services

import (
	"fmt"
	"strconv"

	"bliss.com/tfcatalogue/entities"
	config "bliss.com/tfcatalogue/internal/database"
	"bliss.com/tfcatalogue/internal/helpers"
)

func GetMerhantProducts(requests map[string]string) (DBResponse[entities.Merchant], error) {
	var merchants []entities.Merchant
	var count int64
	var res DBResponse[entities.Merchant]

	limit, err := strconv.Atoi(requests["limit"])
	page, pageErr := strconv.Atoi(requests["page"])

	dbQueryFields := map[string]string{
		"merchantId": "id",
	}

	if err != nil || limit == 0 {
		limit = 25
	}

	if pageErr != nil || page == 0 {
		page = 1
	}

	query := config.Database.Model(&entities.Merchant{})

	if merchantProductId, ok := requests["merchantProductId"]; ok {
		query = query.Joins("JOIN merchant_products ON merchants.id = merchant_products.merchant_id").
			Where("merchant_products.id = ?", merchantProductId).
			Preload("MerchantProducts", "id = ?", merchantProductId)
	} else {
		query = query.Preload("MerchantProducts")
	}

	for key, value := range requests {
		res, ok := dbQueryFields[key]
		if ok {
			condition := fmt.Sprintf("%s = ?", res)
			query = query.Where(condition, value)

		}

	}

	if err := query.Count(&count).Error; err != nil {
		return DBResponse[entities.Merchant]{}, fmt.Errorf("an error occurred querying total records: %w", err)
	}

	result := query.Scopes(
		helpers.Paginate(page, limit),
	).Find(&merchants)
	if result.Error != nil {
		return DBResponse[entities.Merchant]{}, fmt.Errorf("failed to retrieve merchants: %w", result.Error)
	}

	paginationMetadata := helpers.GetPaginationMetadata(page, int(count), limit)

	res.Result = merchants
	res.Metadata.Pages = paginationMetadata
	return res, nil
}
