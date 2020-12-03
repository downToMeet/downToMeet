/* eslint-env jest */

import React from "react";
import { render } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import { Provider } from "react-redux";
import PropTypes from "prop-types";
import store from "../../app/store";
import Profile from "./Profile";

function Wrapper({ children }) {
  return (
    <Provider store={store}>
      <MemoryRouter>{children}</MemoryRouter>
    </Provider>
  );
}

Wrapper.propTypes = {
  children: PropTypes.node.isRequired,
};

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

  const { getByAltText, getByText } = render(
    <Profile
      user={user}
      isMe
      ownedMeetups={[]}
      attendingMeetups={[]}
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
