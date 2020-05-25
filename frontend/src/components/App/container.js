import { connect } from 'react-redux';

import App from './component';

const mapStateToProps = ({ game, player }) => ({ game, player });

const mapDispatchToProps = null;

export default connect(mapStateToProps, mapDispatchToProps)(App);
