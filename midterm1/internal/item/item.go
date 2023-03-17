package item

import "github.com/Tamir1205/midterm1/internal/storage/items"

type ItemDto struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func MapItemToDto(items []items.Item) []ItemDto {
	dto := make([]ItemDto, 0)
	for _, item := range items {
		itemDto := ItemDto{
			ID:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		}
		dto = append(dto, itemDto)
	}
	return dto
}
