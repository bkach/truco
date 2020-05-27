import * as api from 'lib/api';

export const setGames = (games) => ({
  type: 'SET_GAMES',
  payload: { games },
});

export const loadGames = () => (dispatch) => {
  return api.getGames().then((games) => dispatch(setGames(games)));
};
