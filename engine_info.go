package dezge

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/narslan/chess/engine/uci"
)

type EngineInfo struct {
	ID int `json:"id,omitempty"`

	//Name of the engine.
	Name string `json:"name,omitempty"`

	//The path of the executable.
	Path string `json:"path,omitempty"`

	// The ID that is taken from uci command.
	EngineID string `json:"engine_id,omitempty"`

	//The setting from the engine.
	Options map[string]string `json:"options,omitempty"`
	// Timestamps for engine creation & last update.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// EngineService represents a service for managing engines in the database, CRUD-Style.
type EngineInfoService interface {
	FindByID(ctx context.Context, id int) (*EngineInfo, error)
	FindByEngineID(ctx context.Context, engineID string) (*EngineInfo, error)
	Find(ctx context.Context, filter EngineFilter) ([]*EngineInfo, int, error)
	Create(ctx context.Context, engine *EngineInfo) error
	Update(ctx context.Context, id int, upd EngineUpdate) (*EngineInfo, error)
	Delete(ctx context.Context, id int) error
}

// EngineFilter represents a filter used by FindEngines().
type EngineFilter struct {
	// Filtering fields.
	ID       *int    `json:"id"`
	EngineID *string `json:"engine_id"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// EngineUpdate represents a set of fields to update engine info.
type EngineUpdate struct {
	Name *string `json:"name"`
}

func NewEngineInfo(path string) *EngineInfo {
	m := make(map[string]string, 0)
	return &EngineInfo{Path: path, Options: m}
}

func (e *EngineInfo) Validate() error {
	//check if the engine binary exists.
	_, err := os.Stat(e.Path)
	if err != nil {
		return err
	}
	var logger *log.Logger
	eng, err := uci.Run(e.Path, nil, logger)
	if err != nil {
		return err
	}
	defer eng.Quit()

	for k, v := range eng.Options() {
		e.Options[k] = v.String()
	}
	id, ok := e.Options["id"]
	if !ok {
		return fmt.Errorf("engine id is not found")
	}
	e.EngineID = id

	return nil
}
