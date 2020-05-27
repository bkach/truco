import React, { Component } from 'react';

import List from 'components/List';

class GameList extends Component {
  componentDidMount() {
    this.props.loadGames();
  }

  render() {
    const { games, createGame, setGame, deleteGame } = this.props;

    return (
      <List
        title="Games"
        items={games}
        handleAdd={createGame}
        handleDelete={deleteGame}
        handleSelect={setGame}
      />
    );
  }
}

export default GameList;
