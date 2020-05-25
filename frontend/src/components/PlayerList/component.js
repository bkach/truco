import React, { Component } from 'react';

import List from 'components/List';

class PlayerList extends Component {
  handleDelete = (id) => {
    const { deletePlayer, game } = this.props;
    deletePlayer({ gameId: game.id, playerId: id });
  };

  render() {
    const { players, createPlayer, setPlayer } = this.props;

    return (
      <List
        title="Players"
        items={players}
        handleAdd={createPlayer}
        handleDelete={this.handleDelete}
        handleSelect={setPlayer}
      />
    );
  }
}

export default PlayerList;
