package turnbased

import (
	"github.com/strongo/db"
	"time"
	"github.com/strongo/slices"
	"context"
)

const BoardKind = "B"

type BoardEntityBase struct {
	Created          time.Time
	CreatorUserID    string    `datastore:",noindex,omitempty"`
	UserIDs          []string
	UserNames        []string  `datastore:",noindex"`
	UserWins         []int     `datastore:",noindex"`
	UsersMin         int       `datastore:",noindex,omitempty"`
	UsersMax         int       `datastore:",noindex,omitempty"`
	Round            int       `datastore:",noindex,omitempty"`
	Lang             string    `datastore:",noindex,omitempty"`
	TournamentID     string
	TournamentLeft   time.Time `datastore:",noindex,omitempty"`
	TournamentJoined time.Time `datastore:",noindex,omitempty"`
}

func (b *BoardEntityBase) AddUser(userID, userName string) {
	b.UserIDs = append(b.UserIDs, userID)
	b.UserNames = append(b.UserNames, userName)
	b.UserWins = append(b.UserWins, 0)
}

type BoardEntity struct {
	BoardEntityBase
	UserTimes     []time.Time                     `datastore:",noindex"`
	UserMoves     slices.CommaSeparatedValuesList `datastore:",noindex,omitempty"`
	UserWinCounts []int                           `datastore:",noindex"`
	DrawsCount    int                             `datastore:",noindex,omitempty"`
}

type Board struct {
	db.StringID
	*BoardEntity
}

var _ db.EntityHolder = (*Board)(nil)

func (Board) Kind() string {
	return BoardKind
}

func (canvas *Board) SetEntity(v interface{}) {
	canvas.BoardEntity = v.(*BoardEntity)
}

func (canvas Board) Entity() interface{} {
	return canvas.BoardEntity
}

func (canvas Board) NewEntity() interface{} {
	return &BoardEntity{}
}

func GetBoardByID(c context.Context, database db.Database, boardID string) (board Board, err error) {
	board.ID = boardID
	err = database.Get(c, &board)
	return
}
