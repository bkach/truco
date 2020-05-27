const INITIAL_STATE = {};

export default (state = INITIAL_STATE, { type, payload }) => {
  switch (type) {
    case 'SET_GAME':
      return payload.game;

    case 'RESET_GAME':
      return INITIAL_STATE;

    default:
      return state;
  }
};
