import { connect } from 'react-redux';

import { loadGames } from 'state/games/actions';
import { createGame, setGame } from 'state/game/actions';

import GameList from './component';

const mapStateToProps = ({ games }) => ({ games });

const mapDispatchToProps = {
  loadGames,
  createGame,
  setGame,
};

export default connect(mapStateToProps, mapDispatchToProps)(GameList);
