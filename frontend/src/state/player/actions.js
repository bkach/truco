import * as api from 'lib/api';
import { setGame } from 'state/game/actions';

export const setPlayer = (player) => ({
  type: 'SET_PLAYER',
  payload: { player },
});

export const createPlayer = (name) => (dispatch, getState) => {
  const { game } = getState();

  return api
    .createPlayer({
      gameId: game.id,
      name: name || `Player ${game.players.length + 1}`,
    })
    .then(({ player, game }) => {
      dispatch(setPlayer(player));
      dispatch(setGame(game));
    });
};
