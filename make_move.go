package turnbased

import (
	"context"
	"errors"
	"fmt"
	"github.com/strongo/csv"
	"github.com/strongo/dalgo/dal"
	"time"
)

var (
	ErrOldRound     = errors.New("old round")
	ErrUnknownRound = errors.New("unknown round")
)

func MakeMove(c context.Context, now time.Time, database dal.Database, round int, lang, boardID, userID, userName, move string) (board Board, err error) {
	if board, err = GetBoardByID(c, database, boardID); err != nil {
		if dal.IsNotFound(err) {
			err = nil
			// New canvas
			if round != 1 {
				err = fmt.Errorf("round should be 1, got: %v", round)
				return
			}
			board.Data = &BoardEntity{
				BoardEntityBase: BoardEntityBase{
					Lang:      lang,
					Round:     round,
					Created:   now,
					UserIDs:   []string{userID},
					UserNames: []string{userName},
				},
				UserMoves: csv.String(move),
				UserTimes: []time.Time{now},
			}
		}
		return
	}
	if round < board.Data.Round {
		err = fmt.Errorf("%w: record=%v, received=%v", ErrOldRound, board.Data.Round, round)
		return
	} else if round > board.Data.Round {
		err = fmt.Errorf("%w: record=%v, received=%v", ErrUnknownRound, board.Data.Round, round)
		return
	}
	userIDsCount := len(board.Data.UserIDs)
	userMovesCount := len(board.Data.UserMoves.Values())
	if userMovesCount > userIDsCount {
		err = fmt.Errorf("userMovesCount > userIDsCount: %v > %v", userMovesCount, userIDsCount)
		return
	}
	switch userIDsCount {
	case 1:
		if userID == board.Data.UserIDs[0] {
			board.Data.UserMoves = csv.String(move)
		} else {
			board.Data.UserIDs = append(board.Data.UserIDs, userID)
			board.Data.UserNames = append(board.Data.UserNames, userName)
			board.Data.UserMoves = board.Data.UserMoves.Add(move)
			board.Data.UserTimes = append(board.Data.UserTimes, now)
		}
	case 2:
		switch userMovesCount {
		case 0:
			switch userID {
			case board.Data.UserIDs[0]:
				board.Data.UserMoves = csv.String(move)
			case board.Data.UserIDs[1]:
				board.Data.UserMoves = csv.String("," + move)
			default:

			}
		case 1:
			switch userID {
			case board.Data.UserIDs[0]:
				board.Data.UserMoves = board.Data.UserMoves.Set(0, move)
			case board.Data.UserIDs[1]:
				board.Data.UserMoves = board.Data.UserMoves.Add(move)
				board.Data.UserTimes = append(board.Data.UserTimes, now)
			default:
				err = fmt.Errorf("wrong user id=%v", userID)
			}
		case 2:
			switch userID {
			case board.Data.UserIDs[0]:
				board.Data.UserMoves = board.Data.UserMoves.Set(0, move)
			case board.Data.UserIDs[1]:
				board.Data.UserMoves = board.Data.UserMoves.Set(1, move)
			default:
				err = fmt.Errorf("wrong user id=%v", userID)
			}
		}
	}
	return
}

func NextRound(board Board) {
	board.Data.UserMoves = ""
	board.Data.UserTimes = []time.Time{}
	board.Data.Round += 1
}
