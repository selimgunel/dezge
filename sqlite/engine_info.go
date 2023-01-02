package sqlite

import (
	"context"
	"strings"

	"github.com/narslan/dezge"
)

// EngineInfo has an sqlite instance. It accesses and mutates `events` table.
type EngineInfo struct {
	db *DB
}

// It implements dezge.EngineService over sqlite.
var _ dezge.EngineInfoService = &EngineInfo{}

// NewEngine makes a new instance of Engine from database connection.
func NewEngineInfo(db *DB) *EngineInfo {
	return &EngineInfo{db: db}
}

func (e *EngineInfo) FindByID(ctx context.Context, id int) (*dezge.EngineInfo, error) {

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	// Fetch engine.
	engine, err := findEngineInfoByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

// findByID is a helper function to fetch an engine by id.
// Returns ENOTFOUND if engine does not exist.
func findEngineInfoByID(ctx context.Context, tx *Tx, id int) (*dezge.EngineInfo, error) {
	a, _, err := findEngines(ctx, tx, dezge.EngineFilter{ID: &id})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &dezge.Error{Code: dezge.ENOTFOUND, Message: "not found"}
	}
	return a[0], nil
}

func (e *EngineInfo) FindByEngineID(ctx context.Context, engineID string) (*dezge.EngineInfo, error) {

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	// Fetch engine.
	engine, err := findEngineInfoByEngineID(ctx, tx, engineID)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

// findByEngineID is a helper function to fetch an engine by engine id.
// Returns ENOTFOUND if engine does not exist.
func findEngineInfoByEngineID(ctx context.Context, tx *Tx, engineID string) (*dezge.EngineInfo, error) {
	a, _, err := findEngines(ctx, tx, dezge.EngineFilter{EngineID: &engineID})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, dezge.Errorf(dezge.ENOTFOUND, "engine id not found %s", engineID)
	}
	return a[0], nil
}

func findEngines(ctx context.Context, tx *Tx, filter dezge.EngineFilter) (_ []*dezge.EngineInfo, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := filter.EngineID; v != nil {
		where, args = append(where, "engine_id = ?"), append(args, *v)
	}

	// Execute query to fetch user rows.
	rows, err := tx.QueryContext(ctx, `
		SELECT 
		    id,
		    name,
		    path,
		    engine_id,
			created_at,
		    updated_at,
		    COUNT(*) OVER()
		FROM engines
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into Engine objects.
	engines := make([]*dezge.EngineInfo, 0)
	for rows.Next() {

		var e dezge.EngineInfo
		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Path,
			&e.EngineID,
			(*NullTime)(&e.CreatedAt),
			(*NullTime)(&e.UpdatedAt),
			&n,
		); err != nil {
			return nil, 0, err
		}

		engines = append(engines, &e)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return engines, n, nil
}

func (e *EngineInfo) Find(ctx context.Context, filter dezge.EngineFilter) ([]*dezge.EngineInfo, int, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findEngines(ctx, tx, filter)

}

func (e *EngineInfo) Create(ctx context.Context, engine *dezge.EngineInfo) error {

	//Look up if the engine has been already stored.
	_, err := e.FindByEngineID(ctx, engine.EngineID)
	if dezge.ErrorCode(err) != dezge.ENOTFOUND {
		return dezge.Errorf(dezge.EINVALID, "engine exists already")
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = create(ctx, tx, engine)
	if err != nil {
		return err
	}

	return tx.Commit()

}

func (e *EngineInfo) Update(ctx context.Context, id int, upd dezge.EngineUpdate) (*dezge.EngineInfo, error) {
	return nil, nil

}
func (e *EngineInfo) Delete(ctx context.Context, id int) error {
	return nil

}

// createEngine creates an entry for the engine in the database  .
func create(ctx context.Context, tx *Tx, engine *dezge.EngineInfo) error {

	// Set timestamps to the current time.
	engine.CreatedAt = tx.now
	engine.UpdatedAt = engine.CreatedAt

	// Perform basic field validation.
	err := engine.Validate()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, `
		INSERT INTO engines (
			name,
			path,
			engine_id,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?)
	`,
		engine.Name,
		engine.Path,
		engine.EngineID,
		(*NullTime)(&engine.CreatedAt),
		(*NullTime)(&engine.UpdatedAt),
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	engine.ID = int(id)
	return nil
}
