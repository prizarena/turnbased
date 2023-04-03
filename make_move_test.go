package turnbased

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/strongo/dalgo/dal"
	"github.com/strongo/dalgo/mocks4dal"
	"github.com/strongo/slice"
	"testing"
	"time"
)

func TestMakeMove(t *testing.T) {
	testCases := []struct {
		round           int
		userID          string
		userName        string
		move            string
		expectedUserIDs []string
		expectedMoves   []string
	}{
		{1, "u1", "first", "rock", []string{"u1"}, []string{"rock"}},
		{1, "u1", "first", "paper", []string{"u1"}, []string{"paper"}},
		{1, "u2", "second", "scissors", []string{"u1", "u2"}, []string{"paper", "scissors"}},
		{expectedUserIDs: []string{"u1", "u2"}, expectedMoves: []string{}},
		{2, "u2", "second", "rock", []string{"u1", "u2"}, []string{"", "rock"}},
		{2, "u1", "first", "paper", []string{"u1", "u2"}, []string{"paper", "rock"}},
	}

	var board Board

	c := context.Background()

	var err error

	mockCtrl := gomock.NewController(t)
	mockDB := newMockDB(mockCtrl)

	boards := make(map[string]*BoardEntity)

	mockDB.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(c context.Context, record dal.Record) error {
		switch record.Key().Collection() {
		case BoardKind:
			id := fmt.Sprintf("%v", record.Key().ID)
			board := boards[id]
			if board == nil {
				board = record.Data().(*BoardEntity)
				board.Round = 1
				board.UserIDs = []string{"u1"}
			}
			recordBoard := record.Data().(*BoardEntity)
			*recordBoard = *board
		}
		return nil
	}).AnyTimes()
	mockDB.EXPECT().RunReadwriteTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(c context.Context, f func(context.Context, dal.ReadwriteTransaction) error, options ...dal.TransactionOption) error {
		tx := mocks4dal.NewMockReadwriteTransaction(mockCtrl)
		tx.EXPECT().Set(gomock.Any(), gomock.Any()).DoAndReturn(func(c context.Context, record dal.Record) error {
			switch record.Key().Collection() {
			case BoardKind:
				id := fmt.Sprintf("%v", record.Key().ID)
				board := record.Data().(*BoardEntity)
				if len(board.UserIDs) == 0 && board.Round < 2 {
					return fmt.Errorf("board.Round: %d; board.UserIDs: %v", board.Round, board.UserIDs)
				}
				boards[id] = board
			}
			return nil
		})
		return f(c, tx)
	}).AnyTimes()

	for i, testCase := range testCases {
		if testCase.round == 0 {
			NextRound(board)
		} else {
			board, err = MakeMove(c, time.Now(), mockDB, testCase.round, "en-US", "board1", testCase.userID, testCase.userName, testCase.move)
			if err != nil {
				t.Fatalf("case #%d: unexpected error: %v", i+1, err)
			}
			mockDB.RunReadwriteTransaction(c, func(c context.Context, tx dal.ReadwriteTransaction) error {
				return tx.Set(c, board.Record)
			})
		}
		if board.Data == nil {
			t.Fatalf("case #%v: board.Data == nil", i+1)
		}
		if !slice.Equal(board.Data.UserIDs, testCase.expectedUserIDs) {
			t.Fatalf("case #%d: Unexpected UserIDs=%v, expected: %v", i+1, board.Data.UserIDs, testCase.expectedUserIDs)
		}
		//mockDB.Update(c, &board)
	}
}
