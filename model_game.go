package turnbased

import (
	"time"
	"github.com/strongo/db"
)

type GameEntity struct {
	Created      time.Time
	UserIDs      []string
	TournamentID string    `datastore:",omitempty"`
	Strangers    bool      `datastore:",noindex,omitempty"`
	WinnerUserID string    `datastore:",noindex,omitempty"`
	Finished     time.Time `datastore:",omitempty"`
}

type SingleTurnGameEntity struct {
	GameEntity
	UserMoves []string
}

type SingleTurnGame struct {
	db.StringID
	*SingleTurnGameEntity
}