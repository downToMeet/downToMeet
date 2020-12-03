export const UPDATE_USER_DATA = "UPDATE_USER_DATA";
export const CLEAR_USER_DATA = "CLEAR_USER_DATA";

export const updateUserData = (user) => {
  return {
    type: UPDATE_USER_DATA,
    payload: user,
  };
};

export const clearUserData = () => {
  return {
    type: CLEAR_USER_DATA,
  };
};
