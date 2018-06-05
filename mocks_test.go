package turnbased

import (
	"github.com/strongo/db"
	"github.com/strongo/db/mockdb"
	"context"
)

func newMockDB(c context.Context) (mockDB db.Database) {
	return mockdb.NewMockDB(nil, nil)
}
