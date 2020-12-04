/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";
import { Provider } from "react-redux";
import store from "./app/store";
import App from "./App";

test("renders nav bar", () => {
  const { getByRole } = render(
    <Provider store={store}>
      <App />
    </Provider>
  );

  expect(getByRole("link", { name: /DownToMeet/i })).toBeInTheDocument();
  expect(getByRole("link", { name: /New Meetup/i })).toBeInTheDocument();
});
