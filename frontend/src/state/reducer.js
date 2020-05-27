import { combineReducers } from 'redux';

import games from './games/reducer';
import game from './game/reducer';
import player from './player/reducer';

export default combineReducers({ games, game, player });
