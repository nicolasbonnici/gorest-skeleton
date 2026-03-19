package skeleton

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolasbonnici/gorest/crud"
	"github.com/nicolasbonnici/gorest/database"
	"github.com/nicolasbonnici/gorest/processor"
)

type ItemResource struct {
	processor processor.Processor[Item, ItemCreateDTO, ItemUpdateDTO, ItemResponseDTO]
}

func RegisterItemRoutes(app *fiber.App, db database.Database, config *Config) {
	itemCRUD := crud.New[Item](db)
	hooks := NewItemHooks()
	converter := &ItemConverter{}

	fieldMapping := map[string]string{
		"id":          "id",
		"name":        "name",
		"description": "description",
		"user_id":     "user_id",
		"active":      "active",
		"created_at":  "created_at",
		"updated_at":  "updated_at",
	}

	proc := processor.New(processor.ProcessorConfig[Item, ItemCreateDTO, ItemUpdateDTO, ItemResponseDTO]{
		DB:                 db,
		CRUD:               itemCRUD,
		Converter:          converter,
		PaginationLimit:    20,
		PaginationMaxLimit: config.MaxItems,
		FieldMap:           fieldMapping,
		AllowedFields:      []string{"id", "name", "description", "user_id", "active", "created_at", "updated_at"},
	}).
		WithCreateHook(hooks.CreateHook).
		WithUpdateHook(hooks.UpdateHook).
		WithDeleteHook(hooks.DeleteHook).
		WithGetByIDHook(hooks.GetByIDHook).
		WithGetAllHook(hooks.GetAllHook)

	res := &ItemResource{
		processor: proc,
	}

	app.Post("/api/skeleton", res.Create)
	app.Get("/api/skeleton/:id", res.GetByID)
	app.Get("/api/skeleton", res.GetAll)
	app.Put("/api/skeleton/:id", res.Update)
	app.Delete("/api/skeleton/:id", res.Delete)
}

func (r *ItemResource) Create(c *fiber.Ctx) error {
	return r.processor.Create(c)
}

func (r *ItemResource) GetByID(c *fiber.Ctx) error {
	return r.processor.GetByID(c)
}

func (r *ItemResource) GetAll(c *fiber.Ctx) error {
	return r.processor.GetAll(c)
}

func (r *ItemResource) Update(c *fiber.Ctx) error {
	return r.processor.Update(c)
}

func (r *ItemResource) Delete(c *fiber.Ctx) error {
	return r.processor.Delete(c)
}
