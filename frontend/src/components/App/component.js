import React from 'react';

import GameList from 'components/GameList';
import PlayerList from 'components/PlayerList';

const App = ({ game, player }) => {
  if (!game.id) {
    return <GameList />;
  }

  if (!player.id) {
    return <PlayerList players={game.players} />;
  }

  return <p>Game</p>;
};

export default App;
