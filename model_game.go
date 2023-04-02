package turnbased

import (
	"github.com/strongo/dalgo/record"
	"time"
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
	record.WithID[string]
	*SingleTurnPlayEntity
}
