package services

import (
	"fmt"
	"strings"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
)

func GenerateSKU(brand, model string) entities.SKU {
	b := strings.ToUpper(brand[:3])
	m := strings.ToUpper(brand[:3])

	sku := fmt.Sprintf("%s-%s", b, m)

	return entities.SKU(sku)
}
