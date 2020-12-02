import { UPDATE_USER_DATA, CLEAR_USER_DATA } from "./actions";

const initialState = {
  id: "",
  name: "",
};

const userReducer = (state = initialState, action) => {
  switch (action.type) {
    case UPDATE_USER_DATA:
      return {
        ...state,
        id: action.payload.id,
        name: action.payload.name,
      };
    case CLEAR_USER_DATA:
      return {
        ...state,
        id: "",
        name: "",
      };
    default:
      return state;
  }
};

export default userReducer;
