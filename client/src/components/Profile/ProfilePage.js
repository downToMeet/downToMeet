import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Typography } from "@material-ui/core";
import * as fetcher from "../../lib/fetch";
import Profile from "./Profile";

function ProfilePage({ id }) {
  const [loaded, setLoaded] = useState(false);
  const [user, setUser] = useState(null);
  const [isMe, setIsMe] = useState(false);

  useEffect(() => {
    (async () => {
      const { res, resJSON } = await fetcher.getUserData(id);
      if (!res.ok) {
        setLoaded(true);
        return;
      }
      setUser(resJSON);
      setIsMe(id === "me"); // TODO: this check should compare with the redux user state id instead
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
      ownedMeetups={[]}
      attendingMeetups={[]}
      pendingMeetups={[]}
    />
  );
}

ProfilePage.propTypes = {
  id: PropTypes.string.isRequired,
};

export default ProfilePage;
