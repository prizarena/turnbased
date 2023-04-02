package turnbased

import (
	"context"
	"github.com/golang/mock/gomock"
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
	database := newMockDB(mockCtrl)

	for i, testCase := range testCases {
		if testCase.round == 0 {
			NextRound(board)
		} else {
			board, err = MakeMove(c, time.Now(), database, testCase.round, "en-US", "abc", testCase.userID, testCase.userName, testCase.move)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		}
		if board.BoardEntity == nil {
			t.Fatalf("case #%v: board.RevBoardEntity == nil", i+1)
		}
		if !slice.Equal(board.UserIDs, testCase.expectedUserIDs) {
			t.Fatalf("case #%v: Unexpected UserIDs=%v, expected: %v", i+1, board.UserIDs, testCase.expectedUserIDs)
		}
		//database.Update(c, &board)
	}
}
