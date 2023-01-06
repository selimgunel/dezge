package dezge

import (
	"context"
)

type PGN struct {
	ID   string
	Tags map[string]string
}

type PGNService interface {
	FindByID(ctx context.Context, id string) (*PGN, error)
	Find(ctx context.Context, filter PGNFilter) ([]*PGN, string, error)
	Create(ctx context.Context, engine *PGN) error
	Update(ctx context.Context, id int, upd PGNUpdate) (*PGN, error)
	Delete(ctx context.Context, id int) error
}

func NewPGN(pgnSource string) *PGN {

	// moves, tags := pgn.N(pgnSource)
	// tagsMap := make(map[string]string)
	// for _, v := range tags {
	// 	chunks := strings.Split(v, " ")
	// 	tagsMap[chunks[0]] = chunks[1]
	// }
	// return &PGN{Moves: moves, Tags: tagsMap}
	return nil
}

type PGNFilter struct {
}
type PGNUpdate struct {
}

// type LiPGN struct {
// 	ID    int `json:"id"`
// 	Event string
// 	LichessURL
// 	Date
// 	Round
// 	White
// 	Black
// 	Result
// 	WhiteElo
// 	BlackElo
// 	ECO
// 	Opening
// 	TimeControl
// 	UTCDate
// 	UTCTime
// 	Termination
// 	WhiteRatingDiff
// 	BlackRatingDiff
// }
//
