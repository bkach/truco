import qs from 'querystring';
import { fetchX } from 'lib/http';

const API_BASE = process.env.REACT_APP_API_BASE;

export const getGames = () => {
  return fetchX(`${API_BASE}/games`);
};

export const createGame = ({ name }) => {
  return fetchX(`${API_BASE}/createGame?${qs.stringify({ name })}`);
};

export const getGame = (id) => {
  return fetchX(`${API_BASE}/game?${qs.stringify({ game_id: id })}`);
};

export const createPlayer = ({ gameId, name }) => {
  return fetchX(
    `${API_BASE}/createPlayer?${qs.stringify({ game_id: gameId, name })}`
  );
};

export const dealCards = ({ gameId }) => {
  return fetchX(`${API_BASE}/dealCards?${qs.stringify({ game_id: gameId })}`);
};

export const playCard = ({ gameId, playerId, card }) => {
  return fetchX(
    `${API_BASE}/playCard?${qs.stringify({
      game_id: gameId,
      player_id: playerId,
    })}`,
    {
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ Card: card }),
    }
  );
};
