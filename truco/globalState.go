package truco

type Listener struct {
	gameId string
	onChange func()
}

// Global state
var Games []Game
var debugOn = false

var gamesListChangeListeners = map[string]Listener{}
var gameChangeListeners = map[string]Listener{}
var playerListChangeListeners = map[string]Listener{}
