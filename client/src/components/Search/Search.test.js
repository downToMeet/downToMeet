/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import Search from "./Search";
import { Wrapper } from "../test";

import * as fetcher from "../../lib/fetch";
import {
  searchMockRemoteMeetups,
  getMockMeetupAttendees,
} from "../../lib/mock";

jest.mock("../../lib/fetch");

test("renders search bar and location select", () => {
  const screen = render(<Search />, {
    wrapper: Wrapper,
  });

  const typeSelect = screen.getByLabelText(/type/i);
  expect(typeSelect).toBeInTheDocument();
  expect(screen.queryByLabelText(/radius/i)).not.toBeInTheDocument();

  expect(screen.getByText(/in person/i)).toBeInTheDocument();
  userEvent.selectOptions(typeSelect, "in-person");

  expect(screen.getByLabelText(/radius/i)).toBeInTheDocument();
});

test("buttons work", () => {
  const screen = render(<Search />, {
    wrapper: Wrapper,
  });

  const searchButton = screen.getAllByRole("button", { name: /search/i })[1];
  expect(searchButton).toBeInTheDocument();
  const resetButton = screen.getByRole("button", { name: /reset search/i });
  expect(resetButton).toBeInTheDocument();

  userEvent.click(searchButton);

  const typeSelect = screen.getByLabelText(/type/i);
  userEvent.selectOptions(typeSelect, "in-person");

  userEvent.click(searchButton);

  userEvent.click(resetButton);

  userEvent.selectOptions(typeSelect, "in-person");
  const radiusSelect = screen.getByLabelText(/radius/i);
  userEvent.selectOptions(radiusSelect, "5 km");
});

test("render remote meetups", async () => {
  const meetup = {
    id: "20",
    owner: "42",
    title: "lets go cook!",
    location: {
      url: "https://google.com/",
    },
    tags: ["cooking"],
    time: new Date().toISOString(),
  };

  fetcher.searchForRemoteMeetups.mockImplementation(
    searchMockRemoteMeetups([meetup])
  );

  fetcher.getMeetupAttendees.mockImplementation(
    getMockMeetupAttendees(new Map([[meetup.id, meetup]]))
  );

  const screen = render(<Search />, { wrapper: Wrapper });

  const typeSelect = screen.getByLabelText(/type/i);
  userEvent.selectOptions(typeSelect, "remote");

  const radiusSelect = await screen.queryByLabelText(/radius/i);
  expect(radiusSelect).not.toBeInTheDocument();

  const autocomplete = screen.getByRole("textbox", {
    name: /look for tags/i,
  });
  expect(autocomplete).toBeInTheDocument();

  const searchButton = screen.getAllByRole("button", { name: /search/i })[1];
  expect(searchButton).toBeInTheDocument();
  userEvent.click(searchButton);

  const meetupEl = await screen.findByText(/lets go cook/i);
  expect(meetupEl).toBeInTheDocument();
});
