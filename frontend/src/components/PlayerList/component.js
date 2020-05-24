import React, { Component } from 'react';

import List from 'components/List';

class PlayerList extends Component {
  render() {
    const { players, createPlayer, setPlayer } = this.props;

    return (
      <List items={players} handleAdd={createPlayer} handleSelect={setPlayer} />
    );
  }
}

export default PlayerList;
