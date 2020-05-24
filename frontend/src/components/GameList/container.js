import { connect } from 'react-redux';

import { setGame } from 'state/game/actions';

import GameList from './component';

const mapStateToProps = ({ games }) => ({ games });

const mapDispatchToProps = {
  setGame,
};

export default connect(mapStateToProps, mapDispatchToProps)(GameList);
