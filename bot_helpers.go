package turnbased

import (
	"github.com/strongo/bots-framework/core"
	"github.com/pkg/errors"
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

