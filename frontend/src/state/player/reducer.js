const INITIAL_STATE = {};

export default (state = INITIAL_STATE, { type, payload }) => {
  switch (type) {
    case 'SET_PLAYER':
      return payload.player;

    default:
      return state;
  }
};
