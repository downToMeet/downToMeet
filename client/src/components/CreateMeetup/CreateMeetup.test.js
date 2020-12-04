/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import CreateMeetup from "./CreateMeetup";
import { Wrapper } from "../test";

test("renders form", () => {
  const { getByText, getByTestId, getByRole } = render(<CreateMeetup />, {
    wrapper: Wrapper,
  });

  expect(getByText("Title")).toBeInTheDocument();
  expect(getByText("Time")).toBeInTheDocument();
  expect(getByTestId("select-meetup-type")).toBeInTheDocument();

  expect(getByRole("button", { name: /create meetup/i })).toBeInTheDocument();
  userEvent.click(getByRole("button", { name: /create meetup/i }));
});
