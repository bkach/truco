import React, { Component } from 'react';

import * as api from 'lib/api';

class GameList extends Component {
  state = {
    games: [],
    isLoading: false,
  };

  componentDidMount() {
    this.loadGames();
  }

  loadGames = () => {
    this.setState({ isLoading: true });
    api.getGames().then((games) => this.setState({ games, isLoading: false }));
  };

  handleAddGame = () => {
    const { games } = this.state;
    this.setState({ isLoading: true });
    api.createGame({ name: `Game ${games.length + 1}` }).then(this.loadGames);
  };

  render() {
    const { games, isLoading } = this.state;

    if (isLoading) {
      return <p>Loading...</p>;
    }

    return (
      <ul>
        {games.map((game) => (
          <li key={game.id}>
            <button onClick={() => this.props.setGame(game)}>
              {game.name}
            </button>
          </li>
        ))}
        <li>
          <button onClick={this.handleAddGame}>Add game</button>
        </li>
      </ul>
    );
  }
}

export default GameList;
