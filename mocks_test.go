package turnbased

import (
	"github.com/golang/mock/gomock"
	"github.com/strongo/dalgo/dal"
	"github.com/strongo/dalgo/mock_dal"
)

func newMockDB(ctrl *gomock.Controller) (mockDB dal.Database) {
	return mock_dal.NewMockDatabase(ctrl)
}
