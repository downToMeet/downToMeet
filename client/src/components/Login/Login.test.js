/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";

import { Wrapper } from "../test";
import Login from "./Login";

test("renders login page", () => {
  const { getByRole: getByRôle } = render(<Login />, { wrapper: Wrapper });

  expect(getByRôle("link", { name: /Facebook/i })).toBeInTheDocument();
  expect(getByRôle("link", { name: /Google/i })).toBeInTheDocument();
});
