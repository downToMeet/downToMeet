/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";

import Meetup from "./Meetup";
import { Wrapper } from "../test";

test("renders meetup", () => {
  const meetupID = "20";

  render(<Meetup id={meetupID} />, { wrapper: Wrapper });

  // expect(getByText(/Create/)).toBeInTheDocument();
});
