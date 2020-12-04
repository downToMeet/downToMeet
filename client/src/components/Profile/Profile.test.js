/* eslint-env jest */

import React from "react";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import Profile from "./Profile";
import { Wrapper } from "../test";

test("renders own profile", () => {
  const user = {
    id: "100000",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };

  const time = new Date();
  time.setDate(time.getDate() + 3);

  const { queryByText, getByAltText, getByText } = render(
    <Profile
      user={user}
      isMe
      ownedMeetups={[
        {
          title: "yeet snowballs",
          time: time.toString(),
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
        },
      ]}
      attendingMeetups={[
        {
          title: "eat snowballs",
          time: time.toString(),
          location: {
            coordinates: {
              lat: 40.3519265,
              lon: -74.6596334,
            },
            name: "Halo Pub",
          },
          id: "8",
          owner: "100001",
          tags: ["eating", "food"],
        },
      ]}
      pendingMeetups={[]}
    />,
    { wrapper: Wrapper }
  );

  expect(getByText(/Jammie/i)).toBeInTheDocument();
  expect(getByText(/jamie@jamie/)).toBeInTheDocument();
  expect(getByText(/call me on my cell/)).toBeInTheDocument();

  const profilePic = getByAltText(/profile/i);
  expect(profilePic).toBeInTheDocument();
  expect(profilePic).toHaveAttribute("src", user.profilePic);

  expect(getByText(/yeet snowballs/)).toBeInTheDocument();
  expect(getByText(/Halo Pub/)).toBeInTheDocument();
  expect(getByText(/cooking/)).toBeInTheDocument();

  expect(queryByText(/eat snowballs/)).not.toBeInTheDocument();
  expect(queryByText(/food/)).not.toBeInTheDocument();

  userEvent.click(screen.getByText(/Attending Meetups/));

  expect(getByText(/eat snowballs/)).toBeInTheDocument();
  expect(getByText(/food/)).toBeInTheDocument();
});

test("renders other's profile", () => {
  const user = {
    id: "100000",
    name: "Jammie",
    email: "jamie@jamie.jamie",
    connections: ["Google"],
    contactInfo: "call me on my cell ;)",
    profilePic:
      "https://lh3.googleusercontent.com/a-/AOh14GiGYIz1CtbKrixMpG288ooWDWM3DA53RbQhQWkz9g=s96-c",
  };

  const { getByAltText, getByText, queryByText } = render(
    <Profile
      user={user}
      isMe={false}
      ownedMeetups={[]}
      attendingMeetups={[]}
      pendingMeetups={[]}
    />,
    { wrapper: Wrapper }
  );

  expect(getByText(/Jammie/i)).toBeInTheDocument();
  expect(queryByText(/jamie@jamie/)).not.toBeInTheDocument();
  expect(queryByText(/call me on my cell/)).not.toBeInTheDocument();

  const profilePic = getByAltText(/profile/i);
  expect(profilePic).toBeInTheDocument();
  expect(profilePic).toHaveAttribute("src", user.profilePic);
});
