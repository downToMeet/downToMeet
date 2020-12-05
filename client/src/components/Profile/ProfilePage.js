import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Typography } from "@material-ui/core";
import { useSelector } from "react-redux";
import * as fetcher from "../../lib/fetch";
import Profile from "./Profile";

/**
 * Wrapper for [Profile](#profile). Takes care of fetching user and meetup data.
 */
function ProfilePage({ id }) {
  const [loaded, setLoaded] = useState(false);
  const [editing, setEditing] = useState(false);
  const [user, setUser] = useState(null);
  const [isMe, setIsMe] = useState(false);
  const [owned, setOwned] = useState(null);
  const [attending, setAttending] = useState(null);
  const [pending, setPending] = useState(null);
  const userID = useSelector((state) => state.id);

  const [newName, setNewName] = useState(null);
  const [newContact, setNewContact] = useState(null);

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
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
      (async () => {
        const res = await Promise.all(attendingIDs.map(fetcher.getMeetup));
        setAttending(
          res
            .map(({ resJSON }) => resJSON)
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
      (async () => {
        const res = await Promise.all(pendingIDs.map(fetcher.getMeetup));
        setPending(
          res
            .map(({ resJSON }) => resJSON)
            .sort((a, b) => new Date(a.time) - new Date(b.time))
        );
      })(),
    ]);
  }

  async function onSubmit() {
    const { res, resJSON } = await fetcher.patchUserData(
      user.id,
      user,
      newName,
      newContact
    );
    if (res.ok) {
      setUser(resJSON);
      setEditing(false);
    }
  }

  useEffect(async () => {
    (async () => {
      const { res, resJSON } = await fetcher.getUserData(id);
      if (!res.ok) {
        setLoaded(true);
        return;
      }
      setUser(resJSON);
      setNewName(resJSON.name);
      setNewContact(resJSON.contactInfo);
      setIsMe(id === "me" || id === userID);
      setMeetupLists(
        resJSON.ownedMeetups,
        resJSON.attending,
        resJSON.pendingApproval
      );
      setLoaded(true);
    })();
  }, [id, userID]);

  if (!loaded) {
    return <Typography>Loading...</Typography>;
  }

  return (
    <Profile
      user={user}
      editing={editing}
      setEditing={setEditing}
      isMe={isMe}
      ownedMeetups={owned || []}
      attendingMeetups={attending || []}
      pendingMeetups={pending || []}
      newName={null}
      setNewName={setNewName}
      newContact={null}
      setNewContact={setNewContact}
      onSubmit={onSubmit}
    />
  );
}

ProfilePage.propTypes = {
  /** User ID. */
  id: PropTypes.string.isRequired,
};

export default ProfilePage;
