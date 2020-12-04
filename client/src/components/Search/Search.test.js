/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import Search from "./Search";
import { Wrapper } from "../test";

test("renders search bar and meetup", () => {
  const { getByText, getByLabelText, queryByLabelText } = render(<Search />, {
    wrapper: Wrapper,
  });

  const typeSelect = getByLabelText(/type/i);
  expect(typeSelect).toBeInTheDocument();
  expect(queryByLabelText(/radius/i)).not.toBeInTheDocument();

  expect(getByText(/in person/i)).toBeInTheDocument();
  userEvent.selectOptions(typeSelect, "in-person");

  expect(getByLabelText(/radius/i)).toBeInTheDocument();
});
