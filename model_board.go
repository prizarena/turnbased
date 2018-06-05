package turnbased

import (
	"github.com/strongo/db"
	"time"
	"github.com/strongo/slices"
)

const BoardKind = "B"

type BoardEntity struct {
	Created   time.Time
	UserIDs   []string
	UserTimes []time.Time                    `datastore:",noindex"`
	UserMoves slices.CommaSeparatedValuesList `datastore:",noindex"`
	UserWins  []int                          `datastore:",noindex"`
	Lang      string                         `datastore:",noindex"`
	Round     int                            `datastore:",noindex"`
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
