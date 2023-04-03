package turnbased

import (
	"github.com/golang/mock/gomock"
	"github.com/strongo/dalgo/mocks4dal"
)

func newMockDB(ctrl *gomock.Controller) (mockDB *mocks4dal.MockDatabase) {
	return mocks4dal.NewMockDatabase(ctrl)
}
