package skeleton

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/nicolasbonnici/gorest/crud"
	"github.com/nicolasbonnici/gorest/database"

	_ "github.com/nicolasbonnici/gorest/database/sqlite"
)

// newItemTestDB provisions an isolated in-memory SQLite database with the
// skeleton schema. Production runs on Postgres; SQLite keeps the suite
// hermetic (no external service) while exercising the same query-builder
// code path through the crud layer.
func newItemTestDB(t *testing.T) database.Database {
	t.Helper()

	db, err := database.Open("sqlite", "file:"+t.Name()+"?mode=memory&cache=shared")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	// created_at is populated by the database default, mirroring the Postgres
	// migration; the crud layer intentionally omits it from INSERT statements.
	_, err = db.Exec(context.Background(), `
		CREATE TABLE skeleton_items (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			user_id TEXT NOT NULL,
			active BOOLEAN NOT NULL DEFAULT 1,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME
		)`)
	if err != nil {
		t.Fatalf("create schema: %v", err)
	}

	return db
}

func TestItemCRUD(t *testing.T) {
	db := newItemTestDB(t)
	repo := crud.New[Item](db)
	ctx := context.Background()

	item := Item{
		ID:     uuid.New(),
		Name:   "First item",
		UserID: uuid.New(),
		Active: true,
	}

	if err := repo.Create(ctx, item); err != nil {
		t.Fatalf("Create: %v", err)
	}

	got, err := repo.GetByID(ctx, item.ID)
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if got.Name != item.Name {
		t.Errorf("Name = %q, want %q", got.Name, item.Name)
	}

	got.Name = "Renamed item"
	if err := repo.Update(ctx, item.ID, *got); err != nil {
		t.Fatalf("Update: %v", err)
	}

	updated, err := repo.GetByID(ctx, item.ID)
	if err != nil {
		t.Fatalf("GetByID after update: %v", err)
	}
	if updated.Name != "Renamed item" {
		t.Errorf("Name after update = %q, want %q", updated.Name, "Renamed item")
	}

	if err := repo.Delete(ctx, item.ID); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := repo.GetByID(ctx, item.ID); err == nil {
		t.Error("expected error fetching deleted item, got nil")
	}
}
