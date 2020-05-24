import { combineReducers } from 'redux';

import games from './games/reducer';
import game from './game/reducer';

export default combineReducers({ games, game });
