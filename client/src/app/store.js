import { createStore } from "redux";
import userReducer from "../stores/user/reducers";

const store = createStore(userReducer);

export default store;
