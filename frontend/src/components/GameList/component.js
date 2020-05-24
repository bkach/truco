import React, { Component } from 'react';

import * as api from 'lib/api';

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
      <ul>
        {games.map((game) => (
          <li key={game.id}>
            <button onClick={() => setGame(game)}>{game.name}</button>
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
