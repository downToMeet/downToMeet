import React from "react";
import { MemoryRouter } from "react-router-dom";
import { Provider } from "react-redux";
import PropTypes from "prop-types";
import store from "../app/store";

export function Wrapper({ children }) {
  return (
    <Provider store={store}>
      <MemoryRouter>{children}</MemoryRouter>
    </Provider>
  );
}

Wrapper.propTypes = {
  children: PropTypes.node.isRequired,
};

export const lol = 1;
