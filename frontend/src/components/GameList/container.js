import { connect } from 'react-redux';

import { loadGames } from 'state/games/actions';
import { createGame, setGame, deleteGame } from 'state/game/actions';

import GameList from './component';

const mapStateToProps = ({ games }) => ({ games });

const mapDispatchToProps = {
  loadGames,
  createGame,
  setGame,
  deleteGame,
};

export default connect(mapStateToProps, mapDispatchToProps)(GameList);
