package turnbased

import (
	"github.com/prizarena/prizarena-public/prizarena-client-go"
	"github.com/strongo/db"
		"github.com/strongo/bots-framework/core"
)

const LeaveTournamentCommandCode = "leave-tournament"

func LeaveTournamentAction(whc bots.WebhookContext, prizarenaGameID, prizarenaToken string, database db.Database, board Board) (err error) {
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
	if err = database.Update(c, &board); err != nil {
		return
	}
	return
}
