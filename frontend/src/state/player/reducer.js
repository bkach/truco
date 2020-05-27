const INITIAL_STATE = {};

export default (state = INITIAL_STATE, { type, payload }) => {
  switch (type) {
    case 'SET_PLAYER':
      return payload.player;

    case 'RESET_PLAYER':
      return INITIAL_STATE;

    default:
      return state;
  }
};
