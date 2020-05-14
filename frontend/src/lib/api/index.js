import qs from 'querystring';
import { fetchX } from 'lib/http';

const API_BASE = 'https://bkach.com/api';

// games
export const getGames = () => {
  return fetchX(`${API_BASE}/games`);
};

// game id
export const createGame = ({ name }) => {
  return fetchX(`${API_BASE}/createGame?${qs.stringify({ name })}`);
};

// game
export const getGame = (id) => {
  return fetchX(`${API_BASE}/game?${qs.stringify({ game_id: id })}`);
};

// player id
export const createPlayer = ({ gameId, name }) => {
  return fetchX(
    `${API_BASE}/createPlayer?${qs.stringify({ game_id: gameId, name })}`
  );
};

// no response
export const dealCards = ({ gameId }) => {
  return fetchX(`${API_BASE}/dealCards?${qs.stringify({ game_id: gameId })}`);
};

// no response
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
