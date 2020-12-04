import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Typography } from "@material-ui/core";
import { useSelector } from "react-redux";
import * as fetcher from "../../lib/fetch";
import Profile from "./Profile";

function ProfilePage({ id }) {
  const [loaded, setLoaded] = useState(false);
  const [user, setUser] = useState(null);
  const [isMe, setIsMe] = useState(false);
  const [owned, setOwned] = useState(null);
  const [attending, setAttending] = useState(null);
  const [pending, setPending] = useState(null);
  const userID = useSelector((state) => state.id);

  async function setMeetupLists(
    ownedIDs = [],
    attendingIDs = [],
    pendingIDs = []
  ) {
    return Promise.all([
      (async () => {
        const res = await Promise.all(ownedIDs.map(fetcher.getMeetup));
        setOwned(
          res
            .map(({ resJSON }) => resJSON)
            .filter((meetup) => new Date(meetup.time) > new Date())
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
      (async () => {
        const res = await Promise.all(attendingIDs.map(fetcher.getMeetup));
        setAttending(
          res
            .map(({ resJSON }) => resJSON)
            .filter((meetup) => new Date(meetup.time) > new Date())
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
      (async () => {
        const res = await Promise.all(pendingIDs.map(fetcher.getMeetup));
        setPending(
          res
            .map(({ resJSON }) => resJSON)
            .filter((meetup) => new Date(meetup.time) > new Date())
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
    ]);
  }

  useEffect(async () => {
    (async () => {
      const { res, resJSON } = await fetcher.getUserData(id);
      if (!res.ok) {
        setLoaded(true);
        return;
      }
      setUser(resJSON);
      setIsMe(id === "me" || id === userID);
      setMeetupLists(
        resJSON.ownedMeetups,
        resJSON.attending,
        resJSON.pendingApproval
      );
      setLoaded(true);
    })();
  }, [id]);

  if (!loaded) {
    return <Typography>Loading...</Typography>; // TODO: replace with nice loading screen
  }

  return (
    <Profile
      user={user}
      isMe={isMe}
      ownedMeetups={owned || []}
      attendingMeetups={attending || []}
      pendingMeetups={pending || []}
    />
  );
}

ProfilePage.propTypes = {
  id: PropTypes.string.isRequired,
};

export default ProfilePage;
