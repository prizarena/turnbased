package turnbased

import (
	"context"
	"github.com/strongo/csv"
	"github.com/strongo/dalgo/dal"
	"github.com/strongo/dalgo/record"
	"github.com/strongo/slice"
	"time"
)

const BoardKind = "B"

type BoardEntityBase struct {
	Created        time.Time
	CreatorUserID  string `datastore:",noindex,omitempty"`
	UserIDs        []string
	UserNames      []string `datastore:",noindex"`
	UserWins       []int    `datastore:",noindex"`
	UsersMin       int      `datastore:",noindex,omitempty"`
	UsersMax       int      `datastore:",noindex,omitempty"`
	Round          int      `datastore:",noindex,omitempty"`
	Lang           string   `datastore:",noindex,omitempty"`
	TournamentID   string   `datastore:",omitempty"`
	TournamentJson string   `datastore:",noindex,omitempty"`
}

func (v *BoardEntityBase) AddUser(userID, userName string) {
	v.UserIDs = append(v.UserIDs, userID)
	v.UserNames = append(v.UserNames, userName)
	v.UserWins = append(v.UserWins, 0)
}

type BoardEntity struct {
	BoardEntityBase
	UserTimes     []time.Time `datastore:",noindex"`
	UserMoves     csv.String  `datastore:",noindex,omitempty"`
	UserWinCounts []int       `datastore:",noindex"`
	DrawsCount    int         `datastore:",noindex,omitempty"`
}

type Board struct {
	record.WithID[string]
	*BoardEntity
}

func GetBoardByID(c context.Context, database dal.Database, boardID string) (board Board, err error) {
	board.ID = boardID
	err = database.Get(c, board.Record)
	return
}

func (v BoardEntityBase) IsNewUser(userID string) bool {
	return slice.Index(v.UserIDs, userID) < 0
}

func (v BoardEntityBase) GetUserName(userID string) string {
	for i, id := range v.UserIDs {
		if id == userID {
			return v.UserNames[i]
		}
	}
	panic("unknown user id: " + userID)
}
