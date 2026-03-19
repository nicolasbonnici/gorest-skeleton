package skeleton

import (
	"time"

	"github.com/google/uuid"
)

type ItemConverter struct{}

func (c *ItemConverter) CreateDTOToModel(dto ItemCreateDTO) Item {
	return Item{
		ID:          uuid.New(),
		Name:        dto.Name,
		Description: dto.Description,
		Active:      dto.Active,
		CreatedAt:   time.Now(),
	}
}

func (c *ItemConverter) UpdateDTOToModel(dto ItemUpdateDTO) Item {
	item := Item{}
	if dto.Name != nil {
		item.Name = *dto.Name
	}
	if dto.Description != nil {
		item.Description = *dto.Description
	}
	if dto.Active != nil {
		item.Active = *dto.Active
	}
	return item
}

func (c *ItemConverter) ModelToResponseDTO(model Item) ItemResponseDTO {
	return ItemResponseDTO(model)
}

func (c *ItemConverter) ModelsToResponseDTOs(models []Item) []ItemResponseDTO {
	dtos := make([]ItemResponseDTO, len(models))
	for i, model := range models {
		dtos[i] = c.ModelToResponseDTO(model)
	}
	return dtos
}
