import { LOAD_USER_DATA, UPDATE_USER_DATA } from "./actions";

const initialState = {
  attending: [],
  contactInfo: "",
  coordinates: null,
  userID: "",
  interests: [],
  name: "",
  ownedMeetups: [],
  pendingApproval: [],
};

const userReducer = (state = initialState, action) => {
  switch (action.type) {
    case LOAD_USER_DATA:
      return state;
    case UPDATE_USER_DATA:
      return state;
    default:
      return state;
  }
};

export default userReducer;
