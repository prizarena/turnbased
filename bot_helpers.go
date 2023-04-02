package turnbased

import (
	"context"
	"errors"
	"github.com/strongo/bots-framework/core"
	"github.com/strongo/log"
	"github.com/strongo/slice"
)

func GetBoardID(whi bots.WebhookInput, boardID string) (string, error) {
	if boardID == "" {
		boardID = whi.(bots.WebhookCallbackQuery).GetInlineMessageID()
		if boardID == "" {
			return "", errors.New("expecting to get inlineMessageID")
		}
	}
	return boardID, nil
}

type BoardUsersManagers struct {
	addUserToBoardCalled int
}

func (m BoardUsersManagers) AddUserToBoard(
	c context.Context, userID, userName string, boardBase BoardEntityBase,
	getAppUser func() (bots.BotAppUser, error),
) (userName2 string, boardBase2 BoardEntityBase, err error) {
	userName2 = userName
	boardBase2 = boardBase
	if m.addUserToBoardCalled++; m.addUserToBoardCalled > 1 {
		err = errors.New("method BoardUsersManagers.AddUserToBoard() should be called just once")
		return
	}
	log.Debugf(c, "addUserToBoard")
	var botAppUser bots.BotAppUser
	if slice.Index(boardBase2.UserIDs, userID) < 0 {
		if userName == "" {
			if botAppUser, err = getAppUser(); err != nil {
				return
			}
			userName2 = botAppUser.GetFullName()
		}
		boardBase2.AddUser(userID, userName2)
	}
	return
}
