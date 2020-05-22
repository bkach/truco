package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type CreateGameResponse struct {
	Game truco.Game `json:"game"`
}

func CreateGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, r *http.Request) {
		name, err := util.GetQuery(r, "name")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		game, err := truco.CreateGameAndAddToGames(name)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreateGameResponse{
			Game: *game,
		})

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		fmt.Printf("\nCreated game %s", game.Id)
	})
}
