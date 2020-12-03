import { UPDATE_USER_DATA, CLEAR_USER_DATA } from "./actions";

const initialState = Object.freeze({
  id: null,
  name: null,
  profilePic: null,
});

const userReducer = (state = initialState, action) => {
  switch (action.type) {
    case UPDATE_USER_DATA:
      return {
        ...state,
        id: action.payload.id,
        name: action.payload.name,
        profilePic: action.payload.profilePic,
      };
    case CLEAR_USER_DATA:
      return initialState;
    default:
      return state;
  }
};

export default userReducer;
