import { connect } from 'react-redux';

import { createPlayer, setPlayer, deletePlayer } from 'state/player/actions';

import PlayerList from './component';

const mapStateToProps = ({ game }) => ({ game });

const mapDispatchToProps = {
  createPlayer,
  setPlayer,
  deletePlayer,
};

export default connect(mapStateToProps, mapDispatchToProps)(PlayerList);
