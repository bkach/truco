package actions

import (
	"net/http"
	"strings"
)

func (as *ActionSuite) Test_NewGame() {
	debugOn = false

	// Create a new game
	res := as.JSON("/newGame").Get()

	as.Equal(http.StatusOK, res.Code)

	// Check that the game state is as expected
	as.Equal(Game{
		Board:        []Card{},
		PlayerStates: []PlayerState{},
	}, currentGame)

	resultTrimmed := strings.TrimSpace(res.Body.String())

	// Check that the result is as expected
	as.Equal(resultTrimmed, `{"Board":[],"PlayerStates":[]}`)
}

func (as *ActionSuite) Test_AddPlayer() {
	debugOn = false

	as.JSON("/newGame").Get()

	res := as.JSON("/addPlayer").Post(&AddPlayerRequest{Name: "boris"})
	as.Equal(http.StatusOK, res.Code)

	resultTrimmed := strings.TrimSpace(res.Body.String())

	as.Equal(resultTrimmed, `{"PlayerState":{"Info":{"Name":"boris","ID":"player_boris"},"Hand":[{"value":1,"house":"gold"},{"value":1,"house":"cups"},{"value":1,"house":"spades"}]},"Board":[]}`)
}
