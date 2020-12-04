import { createStore } from "redux";
import userReducer from "../stores/user/reducers";

export function makeStore() {
  return createStore(userReducer);
}

export default makeStore();
