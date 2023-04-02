package turnbased

import (
	"context"
	"errors"
	"github.com/prizarena/prizarena-public/prizarena-client-go"
	"github.com/strongo/bots-framework/core"
	"github.com/strongo/dalgo/dal"
)

const LeaveTournamentCommandCode = "leave-tournament"

func LeaveTournamentAction(whc bots.WebhookContext, prizarenaGameID, prizarenaToken string, database dal.Database, board Board) (err error) {
	c := whc.Context()
	httpClient := whc.BotContext().BotHost.GetHTTPClient(c)
	apiClient := prizarena.NewHttpApiClient(httpClient, "", prizarenaGameID, prizarenaToken)
	if err = apiClient.LeaveTournament(c, board.ID); err != nil {
		return
	}
	if board.BoardEntity == nil {
		if board, err = GetBoardByID(c, database, board.ID); err != nil {
			return err
		}
	}
	panic("not implemented")
	// board.TournamentLeft = time.Now()
	database.RunReadwriteTransaction(c, func(c context.Context, tx dal.ReadwriteTransaction) error {
		//if err = tx.Set(c, &board); err != nil {
		//	return
		//}
		return errors.New("not implemented")
	})
	return
}
