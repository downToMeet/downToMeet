export const UPDATE_USER_DATA = "UPDATE_USER_DATA";
export const CLEAR_USER_DATA = "CLEAR_USER_DATA";

export const updateUserData = ({ id, name }) => {
  const loadObj = { id, name };
  return {
    type: UPDATE_USER_DATA,
    payload: loadObj,
  };
};

export const clearUserData = () => {
  return {
    type: CLEAR_USER_DATA,
  };
};
