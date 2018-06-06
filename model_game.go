package turnbased

import (
	"time"
	"github.com/strongo/db"
)

type PlayEntity struct {
	Created      time.Time
	UserIDs      []string
	TournamentID string    `datastore:",omitempty"`
	Strangers    bool      `datastore:",noindex,omitempty"`
	WinnerUserID string    `datastore:",noindex,omitempty"`
	Finished     time.Time `datastore:",omitempty"`
}

type SingleTurnPlayEntity struct {
	PlayEntity
	UserMoves []string
}

type SingleTurnPlay struct {
	db.StringID
	*SingleTurnPlayEntity
}