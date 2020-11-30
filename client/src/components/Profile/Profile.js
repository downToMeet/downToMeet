/* eslint-disable react/jsx-props-no-spreading */
import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { AppBar, Box, Container, Tab, Typography } from "@material-ui/core";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
import { TabList, TabContext, TabPanel } from "@material-ui/lab";
import { makeStyles, withStyles } from "@material-ui/styles";
import { Link } from "react-router-dom";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles(() => ({
  profilePic: {
    borderRadius: "100%",
    margin: "1.5rem",
    maxWidth: "150px",
  },
  meetupInfoTitle: {
    marginTop: "1rem",
    marginBottom: ".5rem",
  },
}));

const StyledTabList = withStyles({
  root: {
    backgroundColor: "white",
  },
  flexContainer: {
    justifyContent: "space-around",
  },
  indicator: {
    backgroundColor: "black",
  },
})((props) => <TabList {...props} />);

const StyledTab = withStyles({
  root: {
    color: "black",
    textTransform: "none",
  },
})((props) => <Tab disableRipple {...props} />);

const StyledTabPanel = withStyles({
  root: {
    backgroundColor: "white",
    width: "100%",
  },
})((props) => <TabPanel {...props} />);

function Profile({ id }) {
  const classes = useStyles();
  const [loaded, setLoaded] = useState(false);
  const [user, setUser] = useState(null);
  const [isMe, setIsMe] = useState(false);
  const [tabValue, setTabValue] = useState("1");

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

  const renderOwnedMeetups = () => {
    // TODO: get actual meetups if they exist
    return (
      <Typography>
        You don’t own any meetups. <Link to="/create">Go create one!</Link>
      </Typography>
    );
  };

  const renderAttendingMeetups = () => {
    // TODO: get actual meetups if they exist
    return (
      <Typography>
        You are not attending any meetups.{" "}
        <Link to="/">Look for one to attend!</Link>
      </Typography>
    );
  };

  const renderPendingMeetups = () => {
    // TODO: get actual meetups if they exist
    return (
      <Typography>
        You don’t have any pending requests.{" "}
        <Link to="/">Look for a meetup to attend!</Link>
      </Typography>
    );
  };

  const renderPersonalInfo = () => {
    if (!isMe) {
      return null;
    }
    return (
      <>
        {user.contactInfo && (
          <Typography>Contact info: {user.contactInfo}</Typography>
        )}
        {user.email && <Typography>Email: {user.email}</Typography>}
        {user.location && (
          <Typography>
            Location: lat {user.location.lat}, lon {user.location.lon}
          </Typography>
          // TODO: replace with sensible representation
        )}
        <Typography
          variant="h4"
          component="h3"
          className={classes.meetupInfoTitle}
        >
          Your meetups:
        </Typography>
        <TabContext value={tabValue}>
          <AppBar position="static">
            <StyledTabList
              onChange={(_, newValue) => setTabValue(newValue)}
              variant="scrollable"
              scrollButtons="auto"
            >
              <StyledTab label="Owned Meetups" value="1" />
              <StyledTab label="Attending Meetups" value="2" />
              <StyledTab label="Pending Meetups" value="3" />
            </StyledTabList>
          </AppBar>
          <StyledTabPanel value="1">{renderOwnedMeetups()}</StyledTabPanel>
          <StyledTabPanel value="2">{renderAttendingMeetups()}</StyledTabPanel>
          <StyledTabPanel value="3">{renderPendingMeetups()}</StyledTabPanel>
        </TabContext>
      </>
    );
  };

  const renderUser = () => {
    return (
      <Container maxWidth="md">
        <Box display="flex" flexDirection="column" alignItems="center">
          {user.profilePic && (
            <img
              src={user.profilePic}
              alt="profile pic"
              className={classes.profilePic}
            />
          )}
          {!user.profilePic && (
            <div
              style={{ position: "relative", width: "150px", height: "150px" }}
            >
              <AccountCircleIcon
                style={{
                  position: "absolute",
                  left: 0,
                  top: 0,
                  width: "100%",
                  height: "100%",
                }}
              />
            </div>
          )}
          <Typography component="h2" variant="h3">
            {user.name}
          </Typography>
          {renderPersonalInfo()}
        </Box>
      </Container>
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
