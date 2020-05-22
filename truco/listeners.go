package truco

func RegisterGameListChangeListener(id string, onChange func()) {
	gamesListChangeListeners[id] = Listener{
		onChange: onChange,
	}
}

func RegisterGameChangeListener(id string, gameId string, onChange func()) {
	gameChangeListeners[id] = Listener{
		gameId: gameId,
		onChange: onChange,
	}
}

func RegisterPlayerListChangeListener(id string, gameId string, onChange func()) {
	playerListChangeListeners[id] = Listener{
		gameId: gameId,
		onChange: onChange,
	}
}

// Removes listeners

func RemoveGameListChangeListener(id string) {
	delete(gamesListChangeListeners, id)
}

func RemoveGameChangeListener(id string) {
	delete(gameChangeListeners, id)
}

func RemovePlayerListChangeListener(id string) {
	delete(playerListChangeListeners, id)
}

// Notifies listener

func notifyGameListChangeListeners() {
	for _, listener := range gamesListChangeListeners {
		listener.onChange()
	}
}

func notifyGameChangeListeners(gameId string) {
	for _, listener := range gameChangeListeners {
		if listener.gameId == gameId {
			listener.onChange()
		}
	}
}

func notifyPlayerListChangeListeners(gameId string) {
	for _, listener := range playerListChangeListeners {
		if listener.gameId == gameId {
			listener.onChange()
		}
	}
}
