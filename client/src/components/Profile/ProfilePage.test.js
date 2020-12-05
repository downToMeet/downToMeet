/* eslint-env jest */

import React from "react";
import { MemoryRouter } from "react-router-dom";
import { Provider } from "react-redux";
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import { makeStore } from "../../app/store";
import { updateUserData } from "../../stores/user/actions";
import ProfilePage from "./ProfilePage";
import * as fetcher from "../../lib/fetch";
import {
  getMockUserData,
  getMockMeetup,
  getMockMeetupAttendees,
} from "../../lib/mock";

jest.mock("../../lib/fetch");

test("renders profile as self", async () => {
  const store = makeStore();

  const me = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date(),
    contactInfo: "TIM-TIM-TIMM",
    email: "tim@tim.tim",
    ownedMeetups: ["20"],
  };

  const meetupOwned = {
    id: "20",
    owner: me.id,
    title: "lets go swimming!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: new Date().toISOString(),
    attendees: [],
    pendingAttendees: [],
  };

  store.dispatch(updateUserData(me));

  fetcher.getUserData.mockImplementation(
    getMockUserData(new Map([[me.id, me]]))
  );

  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetupOwned.id, meetupOwned]]))
  );

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <ProfilePage id={me.id} />
      </MemoryRouter>
    </Provider>
  );

  const email = await screen.findByText(/tim@tim.tim/i);
  expect(email).toBeInTheDocument();
  const contact = await screen.findByText(/TIM-TIM-TIMM/i);
  expect(contact).toBeInTheDocument();
});

test("renders non-existent user", async () => {
  const store = makeStore();

  const me = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date(),
    contactInfo: "TIM-TIM-TIMM",
    email: "tim@tim.tim",
  };
  store.dispatch(updateUserData(me));

  fetcher.getUserData.mockImplementation(
    getMockUserData(new Map([[me.id, me]]))
  );

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <ProfilePage id="50" />
      </MemoryRouter>
    </Provider>
  );

  const user = await screen.findByText(/was not found/);
  expect(user).toBeInTheDocument();
});

test("renders meetups", async () => {
  const store = makeStore();
  const meetupDate = new Date();
  meetupDate.setDate(meetupDate.getDate() + 10);

  const me = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date(),
    contactInfo: "TIM-TIM-TIMM",
    email: "tim@tim.tim",
    ownedMeetups: ["20", "23"],
    attending: ["21", "24"],
    pendingApproval: ["22", "25"],
  };

  const meetupOwned = {
    id: "20",
    owner: me.id,
    title: "lets go swimming!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: meetupDate.toISOString(),
    attendees: [],
    pendingAttendees: [],
  };

  const meetupAttending = {
    id: "21",
    owner: "30",
    title: "lets go biking!",
    location: {
      url: "https://google.com/",
    },
    tags: ["biking"],
    time: meetupDate.toISOString(),
    attendees: [me.id],
    pendingAttendees: [],
  };

  const meetupPending = {
    id: "22",
    owner: "50",
    title: "lets go hiking!",
    location: {
      url: "https://google.com/",
    },
    tags: ["hiking"],
    time: meetupDate.toISOString(),
    attendees: [],
    pendingAttendees: [me.id],
  };

  const meetupDate2 = new Date();
  meetupDate2.setDate(meetupDate2.getDate() + 11);

  const meetupOwned2 = {
    id: "23",
    owner: me.id,
    title: "lets go swimming again!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: meetupDate2.toISOString(),
    attendees: [],
    pendingAttendees: [],
  };

  const meetupAttending2 = {
    id: "24",
    owner: "30",
    title: "lets go biking again!",
    location: {
      url: "https://google.com/",
    },
    tags: ["biking"],
    time: meetupDate2.toISOString(),
    attendees: [me.id],
    pendingAttendees: [],
  };

  const meetupPending2 = {
    id: "25",
    owner: "50",
    title: "lets go hiking again!",
    location: {
      url: "https://google.com/",
    },
    tags: ["hiking"],
    time: meetupDate2.toISOString(),
    attendees: [],
    pendingAttendees: [me.id],
  };

  store.dispatch(updateUserData(me));

  fetcher.getUserData.mockImplementation(
    getMockUserData(new Map([[me.id, me]]))
  );

  const meetupMap = new Map([
    [meetupOwned.id, meetupOwned],
    [meetupAttending.id, meetupAttending],
    [meetupPending.id, meetupPending],
    [meetupOwned2.id, meetupOwned2],
    [meetupAttending2.id, meetupAttending2],
    [meetupPending2.id, meetupPending2],
  ]);

  fetcher.getMeetup.mockImplementation(getMockMeetup(meetupMap));

  fetcher.getMeetupAttendees.mockImplementation(
    getMockMeetupAttendees(meetupMap)
  );

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <ProfilePage id={me.id} />
      </MemoryRouter>
    </Provider>
  );

  const owned = await screen.findByText(/lets go swimming!/i);
  const owned2 = await screen.findByText(/lets go swimming again!/i);
  expect(owned).toBeInTheDocument();
  expect(owned2).toBeInTheDocument();

  userEvent.click(screen.getByText(/Attending Meetups/i));
  const attending = await screen.findByText(/lets go biking!/);
  const attending2 = await screen.findByText(/lets go biking again!/);
  expect(attending).toBeInTheDocument();
  expect(attending2).toBeInTheDocument();

  userEvent.click(screen.getByText(/Pending Meetups/i));
  const pending = await screen.findByText(/lets go hiking!/);
  const pending2 = await screen.findByText(/lets go hiking again!/);
  expect(pending).toBeInTheDocument();
  expect(pending2).toBeInTheDocument();
});
