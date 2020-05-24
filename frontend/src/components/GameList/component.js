import React, { Component } from 'react';

import * as api from 'lib/api';
import List from 'components/List';

class GameList extends Component {
  state = {
    isLoading: false,
  };

  componentDidMount() {
    this.loadGames();
  }

  loadGames = () => {
    this.setState({ isLoading: true });
    api.getGames().then((games) => {
      this.props.setGames(games);
      this.setState({ isLoading: false });
    });
  };

  handleAddGame = () => {
    const { games } = this.state;
    this.setState({ isLoading: true });
    api.createGame({ name: `Game ${games.length + 1}` }).then(this.loadGames);
  };

  render() {
    const { games, setGame } = this.props;
    const { isLoading } = this.state;

    if (isLoading) {
      return <p>Loading...</p>;
    }

    return (
      <List
        items={games}
        handleAdd={this.handleAddGame}
        handleSelect={setGame}
      />
    );
  }
}

export default GameList;
