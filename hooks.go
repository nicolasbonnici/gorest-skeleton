package skeleton

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	auth "github.com/nicolasbonnici/gorest/auth"
	"github.com/nicolasbonnici/gorest/crud"
	"github.com/nicolasbonnici/gorest/query"
)

type ItemHooks struct{}

func NewItemHooks() *ItemHooks {
	return &ItemHooks{}
}

func (h *ItemHooks) CreateHook(c *fiber.Ctx, dto ItemCreateDTO, model *Item) error {
	name := strings.TrimSpace(dto.Name)
	if name == "" {
		return fiber.NewError(400, "name is required")
	}
	if len(name) > 255 {
		return fiber.NewError(400, "name exceeds maximum length of 255 characters")
	}

	if len(dto.Description) > 1000 {
		return fiber.NewError(400, "description exceeds maximum length of 1000 characters")
	}

	model.Name = name

	user := auth.GetAuthenticatedUser(c)
	if user != nil {
		userID, err := uuid.Parse(user.UserID)
		if err == nil {
			model.UserID = userID
		}
	}

	return nil
}

func (h *ItemHooks) UpdateHook(c *fiber.Ctx, dto ItemUpdateDTO, model *Item) error {
	if dto.Name == nil && dto.Description == nil && dto.Active == nil {
		return fiber.NewError(400, "at least one field must be provided")
	}

	if dto.Name != nil {
		name := strings.TrimSpace(*dto.Name)
		if name == "" {
			return fiber.NewError(400, "name cannot be empty")
		}
		if len(name) > 255 {
			return fiber.NewError(400, "name exceeds maximum length of 255 characters")
		}
		model.Name = name
	}

	if dto.Description != nil && len(*dto.Description) > 1000 {
		return fiber.NewError(400, "description exceeds maximum length of 1000 characters")
	}

	return nil
}

func (h *ItemHooks) DeleteHook(c *fiber.Ctx, id any) error {
	return nil
}

func (h *ItemHooks) GetByIDHook(c *fiber.Ctx, id any) error {
	return nil
}

func (h *ItemHooks) GetAllHook(c *fiber.Ctx, conditions *[]query.Condition, orderBy *[]crud.OrderByClause) error {
	return nil
}
