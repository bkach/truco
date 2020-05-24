import { connect } from 'react-redux';

import { setGames } from 'state/games/actions';
import { setGame } from 'state/game/actions';

import GameList from './component';

const mapStateToProps = ({ games }) => ({ games });

const mapDispatchToProps = {
  setGames,
  setGame,
};

export default connect(mapStateToProps, mapDispatchToProps)(GameList);
