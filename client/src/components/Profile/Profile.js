import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Typography } from "@material-ui/core";
import * as fetcher from "../../lib/fetch";

function Profile({ id }) {
  const [loaded, setLoaded] = useState(false);
  const [user, setUser] = useState(null);

  useEffect(() => {
    (async () => {
      const { res, resJSON } = await fetcher.getUserData(id);
      if (!res.ok) {
        setLoaded(true);
        return;
      }
      setUser(resJSON);
      setLoaded(true);
    })();
  }, [id]);

  const renderUser = () => {
    return (
      <Typography>
        Name is {user.name}, ID is {id}
      </Typography>
    );
  };

  const renderNotFound = () => {
    return <Typography>Specified user was not found.</Typography>;
  };

  if (!loaded) {
    return <Typography>Loading...</Typography>; // TODO: replace with nice loading screen
  }
  if (user === null) {
    return renderNotFound();
  }
  return renderUser();
}

Profile.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Profile;
