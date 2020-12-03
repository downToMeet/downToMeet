/* eslint-disable react/jsx-props-no-spreading */
import React, { useState } from "react";
import PropTypes from "prop-types";
import { AppBar, Box, Container, Tab, Typography } from "@material-ui/core";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
import { TabList, TabContext, TabPanel } from "@material-ui/lab";
import { makeStyles, withStyles } from "@material-ui/styles";
import { Link } from "react-router-dom";

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

function Profile({
  user,
  isMe,
  ownedMeetups,
  attendingMeetups,
  pendingMeetups,
}) {
  const classes = useStyles();

  if (user === null) {
    return <Typography>Specified user was not found.</Typography>;
  }

  const [tabValue, setTabValue] = useState("owned");

  // TODO: display meetups.
  let ownedMeetupsEl = (
    <Typography>
      You have some meetups, but we can’t display them yet
    </Typography>
  );
  if (ownedMeetups.length === 0) {
    ownedMeetupsEl = (
      <Typography>
        You don’t own any meetups. <Link to="/create">Go make one!</Link>
      </Typography>
    );
  }

  let attendingMeetupsEl = (
    <Typography>
      You are going to some meetups, but we can’t display them yet
    </Typography>
  );
  if (attendingMeetups.length === 0) {
    attendingMeetupsEl = (
      <Typography>
        You are not attending any meetups.{" "}
        <Link to="/">Look for one to attend!</Link>
      </Typography>
    );
  }

  let pendingMeetupsEl = (
    <Typography>
      You are waiting on some meetups, but we can’t display them yet
    </Typography>
  );
  if (pendingMeetups.length === 0) {
    pendingMeetupsEl = (
      <Typography>
        You don’t have any pending requests.{" "}
        <Link to="/">Look for a meetup to attend!</Link>
      </Typography>
    );
  }

  const personalInfo = isMe && (
    <>
      {user.contactInfo && (
        <Typography>Contact info: {user.contactInfo}</Typography>
      )}
      {user.email && <Typography>Email: {user.email}</Typography>}
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
            indicatorColor="primary"
            variant="scrollable"
            scrollButtons="auto"
          >
            <StyledTab label="Owned Meetups" value="owned" />
            <StyledTab label="Attending Meetups" value="attending" />
            <StyledTab label="Pending Meetups" value="pending" />
          </StyledTabList>
        </AppBar>
        <StyledTabPanel value="owned">{ownedMeetupsEl}</StyledTabPanel>
        <StyledTabPanel value="attending">{attendingMeetupsEl}</StyledTabPanel>
        <StyledTabPanel value="pending">{pendingMeetupsEl}</StyledTabPanel>
      </TabContext>
    </>
  );

  return (
    <Container maxWidth="md">
      <Box display="flex" flexDirection="column" alignItems="center">
        {user.profilePic && (
          <img
            src={user.profilePic}
            alt="profile"
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
        {personalInfo}
      </Box>
    </Container>
  );
}

const userType = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  email: PropTypes.string.isRequired,
  connections: PropTypes.arrayOf(PropTypes.string.isRequired),
  contactInfo: PropTypes.string,
  profilePic: PropTypes.string,
});

const meetupType = PropTypes.shape({
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
});

Profile.propTypes = {
  user: userType,
  isMe: PropTypes.bool.isRequired,
  ownedMeetups: PropTypes.arrayOf(meetupType).isRequired,
  attendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
  pendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
};

Profile.defaultProps = {
  user: null,
};

export default Profile;
