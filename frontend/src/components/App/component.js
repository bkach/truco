import React from 'react';

import GameList from 'components/GameList';
import PlayerList from 'components/PlayerList';

import styles from './styles.module.css';

const Router = ({ game, player }) => {
  if (!game.id) {
    return <GameList />;
  }

  if (!player.id) {
    return <PlayerList players={game.players} />;
  }

  return <p>Game</p>;
};

const App = ({ game, player }) => (
  <div className={styles.root}>
    <Router game={game} player={player} />
  </div>
);

export default App;
