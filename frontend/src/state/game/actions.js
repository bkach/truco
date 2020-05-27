import * as api from 'lib/api';
import { setGames } from 'state/games/actions';

export const setGame = (game) => ({
  type: 'SET_GAME',
  payload: { game },
});

export const resetGame = () => ({
  type: 'RESET_GAME',
});

export const createGame = (name) => (dispatch, getState) => {
  const { games } = getState();

  return api
    .createGame({ name: name || `Game ${games.length + 1}` })
    .then(({ game }) => {
      dispatch(setGames(games.concat(game)));
      dispatch(setGame(game));
    });
};

export const deleteGame = (id) => (dispatch, getState) => {
  const { games, game } = getState();
  return api.deleteGame(id).then(() => {
    if (game.id === id) {
      dispatch(resetGame());
    }

    dispatch(setGames(games.filter((game) => game.id !== id)));
  });
};
