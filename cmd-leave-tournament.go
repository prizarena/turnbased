package turnbased

import (
	"github.com/strongo/bots-framework/core"
	"net/url"
	"github.com/prizarena/prizarena-public/prizarena-client-go"
	"context"
	"github.com/strongo/db"
	"time"
)

const LeaveTournamentCommandCode = "leave-tournament"

func NewLeaveTournamentCommand(prizarenaGameID, prizarenaToken string, database db.Database, onSuccess func(whc bots.WebhookContext, board Board) (m bots.MessageFromBot, err error) ) bots.Command {
	return bots.NewCallbackCommand(LeaveTournamentCommandCode, func(whc bots.WebhookContext, callbackUrl *url.URL) (m bots.MessageFromBot, err error) {
		c := whc.Context()
		var board Board
		board.ID = callbackUrl.Query().Get("board")
		httpClient := whc.BotContext().BotHost.GetHTTPClient(c)
		apiClient := prizarena.NewHttpApiClient(httpClient, "", prizarenaGameID, prizarenaToken)
		if err = apiClient.LeaveTournament(c, board.ID); err != nil {
			return
		}
		err = database.RunInTransaction(c, func(c context.Context) error {
			if board, err = GetBoardByID(c, database, board.ID); err != nil {
				return err
			}
			board.LeftTournament = time.Now()
			err = database.Update(c, &board)
			return err
		}, db.SingleGroupTransaction)
		if err != nil {
			return
		}
		return onSuccess(whc, board)
	})
}

