package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type CreateGameResponse struct {
	GameId string `json:"game_id"`
}

func CreateGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		name, err := queries.Query("name")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		gameId, err := truco.CreateGameAndAddToGames(name)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreateGameResponse{
			GameId: gameId,
		})

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		fmt.Printf("\nCreated game %s", gameId)
	})
}
