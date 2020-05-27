const INITIAL_STATE = [];

export default (state = INITIAL_STATE, { type, payload }) => {
  switch (type) {
    case 'SET_GAMES':
      return payload.games;

    default:
      return state;
  }
};
