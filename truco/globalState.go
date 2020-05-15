package truco

type GameChangeListener struct {
	onGameChanged func()
}

// Global state
var Games []Game
var gameChangeListeners = map[string]GameChangeListener{}
var debugOn = false
