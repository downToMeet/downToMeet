/* eslint-env jest */

import React from "react";
import { MemoryRouter } from "react-router-dom";
import { Provider } from "react-redux";
import { render } from "@testing-library/react";

import { makeStore } from "../../app/store";
import { updateUserData } from "../../stores/user/actions";
import Meetup from "./Meetup";
import * as fetcher from "../../lib/fetch";
import {
  getMockUserData,
  getMockMeetup,
  getMockMeetupAttendees,
} from "../../lib/mock";

jest.mock("../../lib/fetch");

test("renders meetup as owner", async () => {
  const store = makeStore();

  const owner = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date(),
  };
  store.dispatch(updateUserData(owner));

  const pendingAttendee = {
    id: "44",
    name: "Connie",
    profilePic:
      "https://media-exp1.licdn.com/dms/image/C5603AQHbRHQncMSa8g/profile-displayphoto-shrink_400_400/0/1523924666175?e=1612396800&v=beta&t=Lb0P8os8RruOpQp9b7Rcj28teYWslG4Y5hZj1VgYMxc",
    createdAt: new Date().toISOString(),
  };

  const meetup = {
    id: "20",
    owner: owner.id,
    title: "lets go swimming!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: new Date().toISOString(),
    pendingAttendees: [pendingAttendee.id],
  };

  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );

  fetcher.getUserData.mockImplementation(
    getMockUserData(
      new Map([
        [owner.id, owner],
        [pendingAttendee.id, pendingAttendee],
      ])
    )
  );

  fetcher.getMeetupAttendees.mockImplementation(
    getMockMeetupAttendees(new Map([[pendingAttendee.id, pendingAttendee]]))
  );

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <Meetup id={meetup.id} />
      </MemoryRouter>
    </Provider>
  );

  const online = await screen.findByText(/online/i);
  expect(online).toBeInTheDocument();
  expect(online.querySelector("a")).toHaveAttribute(
    "href",
    meetup.location.url
  );
  expect(
    await screen.findByRole("button", { name: /Edit Meetup/i })
  ).toBeInTheDocument();
  expect(await screen.findByText(/connie/i)).toBeInTheDocument();
});

test("renders meetup as attendee", async () => {
  const store = makeStore();

  const owner = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date().toISOString(),
  };

  const attendee = {
    id: "43",
    name: "Jamie",
    profilePic:
      "https://www.jamieliu.me/static/fced3ea12a7975c757cb2dab494f8761/47498/IMG_0162.jpg",
    createdAt: new Date().toISOString(),
  };

  const pendingAttendee = {
    id: "44",
    name: "Connie",
    profilePic:
      "https://media-exp1.licdn.com/dms/image/C5603AQHbRHQncMSa8g/profile-displayphoto-shrink_400_400/0/1523924666175?e=1612396800&v=beta&t=Lb0P8os8RruOpQp9b7Rcj28teYWslG4Y5hZj1VgYMxc",
    createdAt: new Date().toISOString(),
  };

  const meetup = {
    id: "20",
    owner: owner.id,
    title: "lets go swimming!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: new Date().toISOString(),
    attendees: [attendee.id],
    pendingAttendees: [pendingAttendee.id],
  };

  fetcher.getUserData.mockImplementation(
    getMockUserData(
      new Map([
        [owner.id, owner],
        [attendee.id, attendee],
        [pendingAttendee.id, pendingAttendee],
      ])
    )
  );
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );
  fetcher.getMeetupAttendees.mockImplementation(
    getMockMeetupAttendees(new Map([[meetup.id, meetup]]))
  );

  store.dispatch(updateUserData(attendee));

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <Meetup id={meetup.id} />
      </MemoryRouter>
    </Provider>
  );

  const online = await screen.findByText(/online/i);
  expect(online).toBeInTheDocument();
  expect(online.querySelector("a")).toHaveAttribute(
    "href",
    meetup.location.url
  );

  expect(screen.getByAltText(/tim's profile/i)).toHaveAttribute(
    "src",
    owner.profilePic
  );
  expect(screen.getByAltText(/jamie's profile/i)).toHaveAttribute(
    "src",
    attendee.profilePic
  );
  // pending attendees only show up for the owner
  expect(screen.queryByAltText(/connie's profile/i)).not.toBeInTheDocument();

  {
    const leave = await screen.findByRole("button", { name: /leave meetup/i });
    expect(leave).toBeInTheDocument();
    expect(
      screen.queryByRole("button", { name: /edit meetup/i })
    ).not.toBeInTheDocument();
    expect(
      screen.queryByRole("button", { name: /cancel join request/i })
    ).not.toBeInTheDocument();
  }

  store.dispatch(updateUserData(pendingAttendee));
  {
    const cancel = await screen.findByRole("button", {
      name: /cancel join request/i,
    });
    expect(cancel).toBeInTheDocument();
    expect(
      screen.queryByRole("button", { name: /edit meetup/i })
    ).not.toBeInTheDocument();
    expect(
      screen.queryByRole("button", { name: /leave meetup/i })
    ).not.toBeInTheDocument();
  }
});

test("renders meetup as rejected", async () => {
  const store = makeStore();

  const owner = {
    id: "42",
    name: "Tim",
    profilePic:
      "https://avatars1.githubusercontent.com/u/1538624?s=60&u=735bed1f295a88806f5b5b6f033c4eec7fd58fc8&v=4",
    createdAt: new Date().toISOString(),
  };

  const rejectedAttendee = {
    id: "45",
    name: "no one",
    profilePic:
      "https://raw.githubusercontent.com/testing-library/jest-dom/master/other/owl.png",
    createdAt: new Date().toISOString(),
  };

  const meetup = {
    id: "20",
    owner: owner.id,
    title: "lets go swimming!",
    location: {
      url: "https://google.com/",
    },
    tags: ["swimming"],
    time: new Date().toISOString(),
    rejected: true,
  };

  fetcher.getUserData.mockImplementation(
    getMockUserData(
      new Map([
        [owner.id, owner],
        [rejectedAttendee.id, rejectedAttendee],
      ])
    )
  );
  fetcher.getMeetup.mockImplementation(
    getMockMeetup(new Map([[meetup.id, meetup]]))
  );
  fetcher.getMeetupAttendees.mockImplementation(
    getMockMeetupAttendees(new Map([[meetup.id, meetup]]))
  );

  store.dispatch(updateUserData(rejectedAttendee));

  const screen = render(
    <Provider store={store}>
      <MemoryRouter>
        <Meetup id={meetup.id} />
      </MemoryRouter>
    </Provider>
  );

  const online = await screen.findByText(/online/i);
  expect(online).toBeInTheDocument();
  expect(online.querySelector("a")).not.toBeInTheDocument();

  expect(screen.getByAltText(/tim's profile/i)).toHaveAttribute(
    "src",
    owner.profilePic
  );

  expect(
    screen.queryByRole("button", { name: /edit meetup/i })
  ).not.toBeInTheDocument();
  expect(
    screen.queryByRole("button", { name: /cancel join request/i })
  ).not.toBeInTheDocument();
  expect(
    screen.queryByRole("button", { name: /leave meetup/i })
  ).not.toBeInTheDocument();
  expect(screen.queryByText(/you were rejected/i)).toBeInTheDocument();
});
