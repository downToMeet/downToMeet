/* eslint-env jest */

import React from "react";
import { MemoryRouter, Route, Router } from "react-router-dom";
import { Provider } from "react-redux";
import { render, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { createMemoryHistory } from "history";

import { makeStore } from "../../app/store";
import CreateMeetup from "./CreateMeetup";
import { Wrapper } from "../test";
import * as fetcher from "../../lib/fetch";
import { getMockMeetup, mockCreateOrEditMeetup } from "../../lib/mock";
import { updateUserData } from "../../stores/user/actions";

jest.mock("../../lib/fetch");

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

test("renders owned meetup for editing", async () => {
  const store = makeStore();

  // mock data
  const time = new Date();
  time.setDate(time.getDate() + 3);

  const user = {
    id: "100000",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };

  store.dispatch(updateUserData(user));

  const meetup = {
    title: "yeet snowballs",
    time: time.toISOString(),
    location: {
      coordinates: {
        lat: 40.3519265,
        lon: -74.6596334,
      },
      name: "Halo Pub",
    },
    id: "8",
    minCapacity: 2,
    maxCapacity: 8,
    owner: "100000",
    tags: ["cooking", "baking"],
  };

  // mock implementations
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );

  // render
  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <CreateMeetup id={meetup.id} />
      </MemoryRouter>
    </Provider>
  );

  const title = await screen.findByDisplayValue(/yeet snowballs/);
  expect(title).toBeInTheDocument();
});

test("renders owned remote meetup for editing", async () => {
  const store = makeStore();

  // mock data
  const time = new Date();
  time.setDate(time.getDate() + 3);

  const user = {
    id: "100000",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };

  store.dispatch(updateUserData(user));

  const meetup = {
    title: "yeet snowballs",
    time: time.toISOString(),
    location: {
      url: "http://google.com",
    },
    id: "8",
    minCapacity: 2,
    maxCapacity: 8,
    owner: "100000",
    tags: ["cooking", "baking"],
  };

  // mock implementations
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );

  fetcher.createOrEditMeetup.mockImplementation(mockCreateOrEditMeetup());

  const history = createMemoryHistory();
  history.push(`/meetup/${meetup.id}/edit`);

  // render
  const screen = render(
    <Provider store={store}>
      <Router history={history}>
        <Route
          path={history.location.pathname}
          render={() => <CreateMeetup id={meetup.id} />}
        />
      </Router>
    </Provider>
  );

  const title = await screen.findByDisplayValue(/yeet snowballs/);
  expect(title).toBeInTheDocument();
  const submitButton = await screen.findByRole("button", {
    name: /save changes/i,
  });
  expect(submitButton).toBeInTheDocument();
  userEvent.click(submitButton);
  await waitFor(() =>
    expect(history.location.pathname).toBe(`/meetup/${meetup.id}`)
  );
});

test("nonexistent meetup id renders creation page", async () => {
  const store = makeStore();

  // mock data
  const time = new Date();
  time.setDate(time.getDate() + 3);

  const user = {
    id: "100000",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };
  store.dispatch(updateUserData(user));

  const meetup = {
    title: "yeet snowballs",
    time: time.toISOString(),
    location: {
      coordinates: {
        lat: 40.3519265,
        lon: -74.6596334,
      },
      name: "Halo Pub",
    },
    id: "8",
    owner: "100000",
    tags: ["cooking", "baking"],
  };

  // mock implementations
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );

  const history = createMemoryHistory();
  history.push("/meetup/123/edit");

  // render
  render(
    <Provider store={store}>
      <Router history={history}>
        <Route
          path="/meetup/123/edit"
          render={() => <CreateMeetup id="123" />}
        />
      </Router>
    </Provider>
  );

  await waitFor(() => expect(history.location.pathname).toBe("/create"));
});

test("unowned meetup redirects to info page", async () => {
  const store = makeStore();

  // mock data
  const time = new Date();
  time.setDate(time.getDate() + 3);

  const user = {
    id: "100001",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };
  store.dispatch(updateUserData(user));

  const meetup = {
    title: "yeet snowballs",
    time: time.toISOString(),
    location: {
      coordinates: {
        lat: 40.3519265,
        lon: -74.6596334,
      },
      name: "Halo Pub",
    },
    id: "8",
    owner: "100000",
    tags: ["cooking", "baking"],
  };

  // mock implementations
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );

  const history = createMemoryHistory();
  history.push("/meetup/8/edit");

  // render
  render(
    <Provider store={store}>
      <Router history={history}>
        <Route path="/meetup/8/edit" render={() => <CreateMeetup id="8" />} />
      </Router>
    </Provider>
  );

  await waitFor(() => expect(history.location.pathname).toBe("/meetup/8"));
});

test("test validate", async () => {
  const store = makeStore();
  // mock data
  const time = new Date();
  time.setDate(time.getDate() + 3);

  const user = {
    id: "100001",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };
  store.dispatch(updateUserData(user));

  const history = createMemoryHistory();
  history.push("/create");

  const screen = render(
    <Provider store={store}>
      <Router history={history}>
        <Route path="/create" render={() => <CreateMeetup />} />
      </Router>
    </Provider>
  );

  const submitButton = await screen.findByRole("button", {
    name: /create meetup/i,
  });
  expect(submitButton).toBeInTheDocument();
  userEvent.click(submitButton);

  expect(
    await screen.findByText(/Please ensure all required/i)
  ).toBeInTheDocument();
});
