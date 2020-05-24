import { connect } from 'react-redux';

import { createPlayer, setPlayer } from 'state/player/actions';

import PlayerList from './component';

const mapStateToProps = ({ player }) => ({ player });

const mapDispatchToProps = {
  createPlayer,
  setPlayer,
};

export default connect(mapStateToProps, mapDispatchToProps)(PlayerList);
